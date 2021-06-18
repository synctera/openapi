package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

type multiStringFlag []string

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
	var internalApis multiStringFlag

	flag.Var(&projects, "project", "project name and optional reference, use this multiple times for multiple projects, format repo[:branch] (eg mainapi:my_branch)")
	flag.Var(&internalApis, "internal", "API endpoint name to be considered internal only (eg signups)")
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

	isInternal := func(apiRootSpec string) bool {
		entity := entityName(apiRootSpec)
		for _, internalApi := range internalApis {
			if internalApi == entity {
				return true
			}
		}
		return false
	}

	const openapiDirName = "openapi"
	outputFileName := func(inputFileName string) string {
		return path.Join(*outputBasePath, strings.TrimPrefix(inputFileName, openapiDirName))
	}

	gitUrlPrefix := "git@gitlab.com:synctera/"
	if ciToken := os.Getenv("CI_JOB_TOKEN"); len(ciToken) > 0 {
		gitUrlPrefix = fmt.Sprintf("https://gitlab-ci-token:%s@gitlab.com/synctera/", ciToken)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	specFiles := specSearchResults{
		all: make(map[string]struct{}),
	}
	for _, project := range projects {
		projectParts := strings.SplitN(project, ":", 2)
		repoName := projectParts[0]
		referenceName := plumbing.HEAD
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

		head, err := repo.Head()
		if err != nil {
			log.Fatalf("Error finding HEAD for %s: %s", project, err)
		}
		log.Printf("Pulling OpenAPI specs from %s:%s", project, head.String())

		tree, err := repo.Worktree()
		if err != nil {
			log.Fatalf("Error getting working tree: %s", err)
		}

		projectSpecFiles := findSpecs(tree.Filesystem, openapiDirName)
		log.Printf("Found %d root API files (%d individual files)", len(projectSpecFiles.apiRoots), len(projectSpecFiles.all))

		for specFileName := range projectSpecFiles.all {
			inputFile, err := tree.Filesystem.Open(specFileName)
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

		specFiles.join(projectSpecFiles)
	}

	log.Printf("TOTAL %d root API files (%d individual files)", len(specFiles.apiRoots), len(specFiles.all))

	{
		var externalApiRoots []string
		var internalApiRoots []string
		for _, apiRoot := range specFiles.apiRoots {
			bundled := fmt.Sprintf("%s-api-bundled.yml", entityName(apiRoot))

			internalApiRoots = append(internalApiRoots, bundled)
			if isInternal(apiRoot) {
				log.Println("-", apiRoot, "(internal only)")
			} else {
				log.Println("-", apiRoot)
				externalApiRoots = append(externalApiRoots, bundled)
			}
		}

		// write JSON config file for openapi-merge-cli to create combined internal and external specs
		// NB: openapi-merge-cli assumes relative paths so don't include outputBasePath
		writeMergeConfig(internalApiRoots, "internal-api-merged-bundled.yml", path.Join(*outputBasePath, "merge-internal-apis.json"))

		// write JSON config file for openapi-merge-cli to create combined external spec
		writeMergeConfig(externalApiRoots, "external-api-merged-bundled.yml", path.Join(*outputBasePath, "merge-external-apis.json"))
	}

	// bundle all API roots so they are self-contained and can be merged to the combined specs
	for _, apiRoot := range specFiles.apiRoots {
		input := outputFileName(apiRoot)
		bundledOutput := path.Join(*outputBasePath, fmt.Sprintf("%s-api-bundled.yml", entityName(apiRoot)))

		log.Println("bundling", input, "->", bundledOutput)
		command := exec.Command("openapi", "bundle", input, "--ext", "yml", "--output", bundledOutput)

		if err := command.Run(); err != nil {
			log.Fatalf("Error bundling %s: %s", input, err)
		}
	}
}

type specSearchResults struct {
	all      map[string]struct{}
	apiRoots []string
}

func (r *specSearchResults) add(specPath string) {
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
}

func findSpecs(fileSystem billy.Filesystem, dirPath string) specSearchResults {
	results := specSearchResults{
		all: make(map[string]struct{}),
	}

	dirEntries, err := fileSystem.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory %s: %s", dirPath, err)
	}

	for _, dirEntry := range dirEntries {
		name := dirEntry.Name()
		fullPath := fileSystem.Join(dirPath, name)
		if dirEntry.Mode().IsDir() {
			// recursive traversal unless name is "gen" (indicating the dir contains generated/bundled spec files)
			if dirEntry.Name() == "gen" {
				continue
			}
			results.join(findSpecs(fileSystem, fullPath))
		} else if dirEntry.Mode().IsRegular() {
			// append full path to results if name ends with .yml or .yaml and name is not openapi.yml
			if !strings.HasSuffix(name, ".yml") && !strings.HasSuffix(name, ".yaml") {
				continue // not a YAML file name
			}
			if name == "openapi.yml" {
				continue // top level openapi.yml is a generated file of all specs combined (for kin router ingestion)
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
