package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"gopkg.in/yaml.v3"
)

const (
	openapiDirName     = "openapi"
	cicdConfigFilename = "cicd-config.yml"
)

type multiStringFlag []string

type cicdConfig struct {
	ExternalDomains []string `yaml:"external_domains"`
}

func (c *cicdConfig) append(fileSystem billy.Filesystem, filename string) {
	f, err := fileSystem.Open(filename)
	if err != nil {
		log.Fatalln("failed to open cicd config " + filename)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("failed to read cicd config " + filename)
	}
	config := cicdConfig{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("bad cicd config yaml %s: %v", filename, err)
	}
	if config.ExternalDomains == nil {
		return
	}
	c.ExternalDomains = append(c.ExternalDomains, config.ExternalDomains...)
}

func (flag multiStringFlag) String() string {
	return strings.Join(flag, ",")
}

func (flag *multiStringFlag) Set(value string) error {
	*flag = append(*flag, value)
	return nil
}

func main() {
	outputBasePath := flag.String("output", "", "Base output path for OpenAPI specification files")
	var projects multiStringFlag
	var externalApis multiStringFlag
	var skipDir multiStringFlag

	flag.Var(&projects, "project", "project name and optional reference, use this multiple times for multiple projects, format repo[:branch] (eg mainapi:my_branch)")
	flag.Var(&externalApis, "external", "API endpoint name to be considered external or public (eg customers)")
	flag.Var(&skipDir, "skip", "List of directory that should be skipped in the merge on each repositories")
	flag.Parse()

	if outputBasePath == nil || len(*outputBasePath) == 0 {
		log.Fatalln("Must specify \"output\" parameter.")
	}
	if len(projects) == 0 {
		log.Fatalln("Must specify at least one \"project\" parameter.")
	}

	entityName := func(apiRootSpec string) string {
		return path.Base(path.Dir(apiRootSpec))
	}

	isInternal := func(versionCICDConfig cicdConfig, apiRootSpec string) bool {
		// TODO we should drop the flag and let each repo to define the external domains
		// this this check should be removed
		entity := entityName(apiRootSpec)
		for _, externalApi := range externalApis {
			if externalApi == entity {
				return false
			}
		}

		for _, externalApi := range versionCICDConfig.ExternalDomains {
			if externalApi == entity {
				return false
			}
		}
		return true
	}

	outputFileName := func(inputFileName string) string {
		return filepath.Join(*outputBasePath, strings.TrimPrefix(inputFileName, openapiDirName))
	}

	gitUrlPrefix := "git@gitlab.com:synctera/"
	if ciToken := os.Getenv("CI_JOB_TOKEN"); len(ciToken) > 0 {
		gitUrlPrefix = fmt.Sprintf("https://gitlab-ci-token:%s@gitlab.com/synctera/", ciToken)
	} else if ciToken := os.Getenv("GITLAB_TOKEN"); len(ciToken) > 0 {
		gitUrlPrefix = fmt.Sprintf("https://gitlab-ci-token:%s@gitlab.com/synctera/", ciToken)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	specFiles := make(map[string]specSearchResults)
	for _, project := range projects {
		projectParts := strings.SplitN(project, ":", 2)
		repoName := projectParts[0]
		referenceName := plumbing.NewBranchReferenceName("main")
		if len(projectParts) > 1 && len(projectParts[1]) > 0 {
			referenceName = plumbing.NewBranchReferenceName(projectParts[1])
		}
		repo, err := git.CloneContext(ctx, memory.NewStorage(), memfs.New(), &git.CloneOptions{
			URL:           gitUrlPrefix + repoName,
			ReferenceName: referenceName,
			SingleBranch:  true,
			Depth:         1,
		})
		if err != nil {
			log.Fatalf("Error cloning %s: %s", project, err)
		}

		tree, err := repo.Worktree()
		if err != nil {
			log.Fatalf("Error getting working tree: %s", err)
		}

		versionRoots := findVersionRoots(tree.Filesystem, openapiDirName)

		// TODO: if both root and v0 are defined to have spec, we use v0 to override
		// Once all repo moved to v0/, then we should change this code
		if versionRoots[openapiDirName] && versionRoots[filepath.Join(openapiDirName, "v0")] {
			log.Printf("WARNING: ignoring spec under openapi root directory since v0/ defined")
			delete(versionRoots, openapiDirName)
		}

		for versionRoot := range versionRoots {
			skipDirMap := make(map[string]bool)
			for _, d := range skipDir {
				skipDirMap[filepath.Join(versionRoot, d)] = true
			}

			projectSpecFiles := findSpecs(tree.Filesystem, versionRoot, versionRoots, skipDirMap)

			// TODO- remove once we drop root spec
			// Root spec will be downloaded to v0/ instead
			if versionRoot == openapiDirName {
				projectSpecFiles.handleRootSpec()
			}

			log.Printf("Found %d root API files (%d individual files) for project %s dir %s",
				len(projectSpecFiles.apiRoots), len(projectSpecFiles.all), project, versionRoot)

			for specFileName := range projectSpecFiles.all {
				// TODO - this is super annoying
				// Remove once we drop root dir. if spec on openapi/, we need to make copy to a form openapi/v0/
				inputFileName := specFileName
				if versionRoot == openapiDirName {
					p := strings.Split(specFileName, string(filepath.Separator))
					pArr := []string{p[0]}
					pArr = append(pArr, p[2:]...)
					inputFileName = filepath.Join(pArr...)
				}

				inputFile, err := tree.Filesystem.Open(inputFileName)
				if err != nil {
					log.Fatalf("Error opening %s: %s", specFileName, err)
				}

				output := outputFileName(specFileName)
				if err := os.MkdirAll(path.Dir(output), os.ModePerm); err != nil {
					log.Fatalf("Error creating directory for %s: %s", output, err)
				}
				outputFile, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
				if err != nil {
					log.Fatalf("Error creating %s: %s", output, err)
				}

				if _, err := io.Copy(outputFile, inputFile); err != nil {
					log.Fatalf("Error copying from %s to %s: %s", specFileName, output, err)
				}
			}

			if versionRoot == openapiDirName {
				versionRoot = filepath.Join(openapiDirName, "v0")
			}
			if v, ok := specFiles[versionRoot]; ok {
				v.join(projectSpecFiles)
				specFiles[versionRoot] = v
			} else {
				specFiles[versionRoot] = projectSpecFiles
			}
		}
	}

	{
		for versionRoot, curFiles := range specFiles {
			version := filepath.Base(versionRoot)
			runCmd([]string{"cp", "-r", "common", filepath.Join(filepath.Join(*outputBasePath), version)})

			var externalApiRoots []string
			var internalApiRoots []string
			for _, apiRoot := range curFiles.apiRoots {
				bundled := fmt.Sprintf("%s-api-bundled.yml", entityName(apiRoot))

				internalApiRoots = append(internalApiRoots, bundled)
				if isInternal(curFiles.cicdConfig, apiRoot) {
					log.Printf("internal only for %s %s", version, apiRoot)
				} else {
					log.Printf("external spec for %s %s", version, apiRoot)
					externalApiRoots = append(externalApiRoots, bundled)
				}
			}

			// write JSON config file for openapi-merge-cli to create combined internal and external specs
			// NB: openapi-merge-cli assumes relative paths so don't include outputBasePath
			writeMergeConfig(internalApiRoots, "internal-api-merged-bundled.yml", path.Join(*outputBasePath, version, "merge-internal-apis.json"))

			// write JSON config file for openapi-merge-cli to create combined external spec
			writeMergeConfig(externalApiRoots, "external-api-merged-bundled.yml", path.Join(*outputBasePath, version, "merge-external-apis.json"))
		}
	}

	// bundle all API roots so they are self-contained and can be merged to the combined specs
	for versionRoot, curFiles := range specFiles {
		version := filepath.Base(versionRoot)
		for _, apiRoot := range curFiles.apiRoots {
			input := outputFileName(apiRoot)
			bundledOutput := path.Join(*outputBasePath, version, fmt.Sprintf("%s-api-bundled.yml", entityName(apiRoot)))

			log.Println("bundling", input, "->", bundledOutput)
			runCmd([]string{"openapi", "bundle", input, "--ext", "yml", "--output", bundledOutput})
		}
	}
}

func findVersionRoots(fileSystem billy.Filesystem, dirPath string) map[string]bool {
	dirEntries, err := fileSystem.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory %s: %s", dirPath, err)
	}

	res := map[string]bool{}
	for _, dirEntry := range dirEntries {
		name := dirEntry.Name()
		fullPath := fileSystem.Join(dirPath, name)
		if dirEntry.Mode().IsDir() {
			// recursive traversal unless name is "gen" (indicating the dir contains generated/bundled spec files)
			if dirEntry.Name() == "gen" {
				continue
			}
			for p := range findVersionRoots(fileSystem, fullPath) {
				res[p] = true
			}
		} else if dirEntry.Mode().IsRegular() && name == "openapi.yml" {
			res[dirPath] = true
		}
	}
	return res
}

type specSearchResults struct {
	all        map[string]struct{}
	apiRoots   []string
	cicdConfig cicdConfig
}

// TODO - can be removed when we drop root dir spec
// This function handles our old fashion the spec is in root, not versione directory
// It considered as v0/ when we copy over
func (r *specSearchResults) handleRootSpec() {
	insertV0 := func(in string) string {
		p := strings.Split(in, string(filepath.Separator))
		newP := make([]string, 0, len(p)+1)
		newP = append(newP, p[0])
		newP = append(newP, "v0")
		newP = append(newP, p[1:]...)
		return filepath.Join(newP...)
	}
	newAll := make(map[string]struct{})
	for k := range r.all {
		newAll[insertV0(k)] = struct{}{}
	}
	r.all = newAll
	for i := range r.apiRoots {
		r.apiRoots[i] = insertV0(r.apiRoots[i])
	}
}

func (r *specSearchResults) add(specPath string) {
	if r.all == nil {
		r.all = make(map[string]struct{})
	}
	if _, exists := r.all[specPath]; exists {
		log.Fatalf("File %s already exists!", specPath)
	}
	r.all[specPath] = struct{}{}

	if path.Base(specPath) == "api.yml" || path.Base(specPath) == "api.yaml" {
		r.apiRoots = append(r.apiRoots, specPath)
	}
}

func (r *specSearchResults) join(other specSearchResults) {
	for spec := range other.all {
		r.add(spec)
	}
	r.cicdConfig.ExternalDomains = append(r.cicdConfig.ExternalDomains, other.cicdConfig.ExternalDomains...)
}

func findSpecs(fileSystem billy.Filesystem, dirPath string, versionRoots, skipDirMap map[string]bool) specSearchResults {
	results := specSearchResults{
		all:        make(map[string]struct{}),
		apiRoots:   []string{},
		cicdConfig: cicdConfig{ExternalDomains: []string{}},
	}

	dirEntries, err := fileSystem.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory %s: %s", dirPath, err)
	}

	for _, dirEntry := range dirEntries {
		name := dirEntry.Name()
		fullPath := fileSystem.Join(dirPath, name)
		if skipDirMap[fullPath] {
			continue
		}
		if dirEntry.Mode().IsDir() {
			// openapi root search that contains v0, v1 etc should be skipped
			if versionRoots[fullPath] {
				continue
			}

			// recursive traversal unless name is "gen" (indicating the dir contains generated/bundled spec files)
			if dirEntry.Name() == "gen" {
				continue
			}
			results.join(findSpecs(fileSystem, fullPath, versionRoots, skipDirMap))
		} else if dirEntry.Mode().IsRegular() {
			// append full path to results if name ends with .yml or .yaml and name is not openapi.yml
			if !strings.HasSuffix(name, ".yml") && !strings.HasSuffix(name, ".yaml") {
				continue // not a YAML file name
			}
			if name == "openapi.yml" {
				continue // top level openapi.yml is a generated file of all specs combined (for kin router ingestion)
			}
			// We cannot keep it in results.all because of the file name conflict, since all services have the same
			// file names. So, we need to load and aggregate from each file and keep the content, not the file name.
			if name == cicdConfigFilename {
				results.cicdConfig.append(fileSystem, fullPath)
				continue
			}
			results.add(fullPath)
		}
	}
	return results
}

// MergeConfig is compatible with openapi-merge-cli
type MergeConfig struct {
	Inputs []MergeConfigInput `json:"inputs"`
	Output string             `json:"output"`
}

type MergeConfigInput struct {
	InputFile string `json:"inputFile"`
}

func writeMergeConfig(inputs []string, output string, config string) {
	configFile, err := os.Create(config)
	if err != nil {
		log.Fatalf("Error creating %s: %s", config, err)
	}

	sort.Strings(inputs) // for consistent generated output

	var mergeConfig MergeConfig
	for _, input := range inputs {
		mergeConfig.Inputs = append(mergeConfig.Inputs, MergeConfigInput{InputFile: input})
	}
	mergeConfig.Output = output

	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(mergeConfig); err != nil {
		log.Fatalf("Error encoding %s: %s", config, err)
	}
}

func runCmd(cmd []string) {
	command := exec.Command(cmd[0], cmd[1:]...)
	var bundleLog bytes.Buffer
	command.Stderr = &bundleLog
	command.Stdout = &bundleLog
	if err := command.Run(); err != nil {
		log.Fatalf("cmd error %v: %s", cmd, bundleLog.String())
	}
}
