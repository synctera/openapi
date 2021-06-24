package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/jdxcode/netrc"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	outputPath := flag.String("output", "", "Output path for common OpenAPI specification files")
	flag.Parse()

	if outputPath == nil || len(*outputPath) == 0 {
		log.Fatalln("Must specify \"output\" parameter.")
	}

	outputDirEntries, err := os.ReadDir(*outputPath)
	if err != nil {
		log.Fatalf("Error opening output directory %s: %s", *outputPath, err)
	}
	for _, dirEntry := range outputDirEntries {
		if !dirEntry.Type().IsRegular() {
			continue
		}

		name := dirEntry.Name()
		if !isYamlFileName(name) {
			continue
		}

		fullPath := path.Join(*outputPath, name)
		log.Println("Removing", fullPath)
		if err := os.Remove(fullPath); err != nil {
			log.Fatalf("Error removing %s: %s", name, err)
		}
	}

	gitlabPassword := ""
	if ciToken := os.Getenv("CI_JOB_TOKEN"); len(ciToken) > 0 {
		gitlabPassword = ciToken
	}

	if gitlabPassword == "" {
		netrcFilePath := os.Getenv("NETRC")
		if netrcFilePath == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Fatalf("Error determining home directory: %s", err)
			}
			netrcFilePath = path.Join(homeDir, ".netrc")
		}

		netrcFile, err := netrc.Parse(netrcFilePath)
		if err != nil {
			log.Fatalf("Error parsing %s: %s", netrcFilePath, err)
		}

		gitlabPassword = netrcFile.Machine("gitlab.com").Get("password")
		if gitlabPassword == "" {
			log.Fatalf("No password for gitlab.com in %s", netrcFilePath)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	repo, err := git.CloneContext(ctx, memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL:           fmt.Sprintf("https://git:%s@gitlab.com/synctera/openapi.git/", gitlabPassword),
		ReferenceName: plumbing.HEAD,
		SingleBranch:  true,
		Depth:         1,
	})
	if err != nil {
		log.Fatalf("Error cloning openapi: %s", err)
	}

	tree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Error getting working tree: %s", err)
	}

	const dirPath = "common"
	fileSystem := tree.Filesystem
	dirEntries, err := fileSystem.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory %s: %s", dirPath, err)
	}

	for _, dirEntry := range dirEntries {
		if !dirEntry.Mode().IsRegular() {
			continue
		}

		name := dirEntry.Name()
		if !isYamlFileName(name) {
			continue
		}

		log.Println("Copying", name)

		fullPath := fileSystem.Join(dirPath, name)
		inputFile, err := fileSystem.Open(fullPath)
		if err != nil {
			log.Fatalf("Error opening %s: %s", fullPath, err)
		}

		output := path.Join(*outputPath, name)
		if err := os.MkdirAll(path.Dir(output), os.ModePerm); err != nil {
			log.Fatalf("Error creating directory for %s: %s", output, err)
		}
		outputFile, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			log.Fatalf("Error creating %s: %s", output, err)
		}

		if _, err := io.Copy(outputFile, inputFile); err != nil {
			log.Fatalf("Error copying from %s to %s: %s", fullPath, output, err)
		}
	}

}

func isYamlFileName(name string) bool {
	return strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".yaml")
}
