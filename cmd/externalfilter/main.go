package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

const externalKey = "x-external"

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("LOGFORMAT") != "json" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05.999",
		})
	}

	var dryRun, debug bool

	pflag.BoolVarP(&dryRun, "dry-run", "n", false, "whether to print filtering changes without performing them")
	pflag.BoolVarP(&debug, "debug", "v", false, "whether to print filtering changes")
	pflag.Parse()

	if pflag.NArg() < 1 {
		log.Fatal().Msgf("Usage: %s <file1> [file2] ...", os.Args[0])
	}

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	for _, file := range pflag.Args() {
		log.Debug().Msgf("Filtering %s", file)
		if err := filterExternal(file, dryRun); err != nil {
			log.Fatal().Err(err).Msgf("Failed to filter %s", file)
		}
	}
}

func filterExternal(path string, dryRun bool) (err error) {
	log := log.With().Str("file", path).Logger()

	msg := func() *zerolog.Event {
		if dryRun {
			return log.Info()
		}

		return log.Debug()
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	var spec map[string]interface{}

	if err := yaml.NewDecoder(file).Decode(&spec); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	paths, ok := spec["paths"].(map[interface{}]interface{})
	if !ok {
		return errors.Errorf("paths was of unexpected type %T", spec["paths"])
	}

	methodKeys := []string{"connect", "delete", "get", "head", "options", "patch", "post", "put", "trace"}

	for pathKey, path := range paths {
		pathMap, ok := path.(map[interface{}]interface{})
		if !ok {
			return errors.Errorf("paths.%s was of unexpected type %T", pathKey, path)
		}

		var found, deleted int

		for _, key := range methodKeys {
			if _, ok := pathMap[key]; !ok {
				continue
			}

			found++

			method, ok := pathMap[key].(map[interface{}]interface{})
			if !ok {
				return errors.Errorf("paths.%s.%s was of unexpected type %T", pathKey, key, pathMap[key])
			}

			if _, ok := method[externalKey]; !ok {
				msg().Msgf("Deleting path %s.%s due to absent %s", pathKey, key, externalKey)

				deleted++
				delete(pathMap, key)
				continue
			}

			external, ok := method[externalKey].(bool)
			if !ok {
				return errors.Errorf("paths.%s.%s.%s was of unexpected type %T", pathKey, key, externalKey, method[externalKey])
			}

			if !external {
				msg().Msgf("Deleting path %s.%s due to %s: %v", pathKey, key, externalKey, external)

				deleted++
				delete(pathMap, key)
				continue
			}
		}

		if found == deleted {
			msg().Msgf("Deleting path %s as it has no remaining methods", pathKey)
			delete(paths, pathKey)
			continue
		} else {
			paths[pathKey] = pathMap
		}
	}

	spec["paths"] = paths
	data, err := yaml.Marshal(spec)
	if err != nil {
		return err
	}

	if dryRun {
		return nil
	}

	return os.WriteFile(path, data, 0644)
}
