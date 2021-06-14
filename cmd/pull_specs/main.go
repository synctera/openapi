package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	outputBasePath := flag.String("output", "", "Base output path for OpenAPI specification files")
	flag.Parse()
	if outputBasePath == nil || len(*outputBasePath) == 0 {
		fmt.Println("Must specify \"output\" parameter.")
		os.Exit(1)
	}


	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	for _, project := range []string{"mainapi"} {
		repo, err := git.CloneContext(ctx, memory.NewStorage(), memfs.New(), &git.CloneOptions{
			URL:          "git@gitlab.com:synctera/" + project,
			SingleBranch: true,
			Depth:        1,
		})
		if err != nil {
			log.Fatalf("Error cloning %s: %s", project, err)
		}

		head, err := repo.Head()
		if err != nil {
			log.Fatalf("Error finding HEAD for %s: %s", project, err)
		}
		log.Printf("Pulling OpenAPI specs from %s %s", project, head.String())

		tree, err := repo.Worktree()
		if err != nil {
			log.Fatalf("Error getting working tree: %s", err)
		}

		const openapiDirName = "openapi"
		specFileNames := findSpecs(tree.Filesystem, openapiDirName)

		for _, specFileName := range specFileNames {
			input, err := tree.Filesystem.Open(specFileName)
			if err != nil {
				log.Fatalf("Error opening %s: %s", specFileName, err)
			}

			outputFileName := path.Join(*outputBasePath, strings.TrimPrefix(specFileName, openapiDirName))
			os.MkdirAll(path.Dir(outputFileName), os.ModePerm)

			output, err := os.Create(outputFileName)
			if err != nil {
				log.Fatalf("Error creating %s: %s", outputFileName, err)
			}
			if _, err := io.Copy(output, input); err != nil {
				log.Fatalf("Error copying from %s to %s: %s", specFileName, outputFileName, err)
			}
		}
	}
}

func findSpecs(fileSystem billy.Filesystem, dirPath string) []string {
	var specFileNames []string

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
			specFileNames = append(specFileNames, findSpecs(fileSystem, fullPath)...)
		} else if dirEntry.Mode().IsRegular() {
			// append full path to results if name ends with .yml or .yaml and name is not openapi.yml
			if !strings.HasSuffix(name, ".yml") && !strings.HasSuffix(name, ".yaml") {
				continue // not a YAML file name
			}
			if name == "openapi.yml" {
				continue // top level openapi.yml is a generated file of all specs combined (for kin router ingestion)
			}
			specFileNames = append(specFileNames, fullPath)
		}
	}

	return specFileNames
}
