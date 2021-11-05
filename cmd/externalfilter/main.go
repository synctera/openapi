package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

const (
	externalKey   = "x-external"
	inKey         = "in"
	refKey        = "$ref"
	nameKey       = "name"
	componentsKey = "components"
	parametersKey = "parameters"

	queryValue = "query"
)

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

func isExternal(node map[interface{}]interface{}) (bool, error) {
	if _, ok := node[externalKey]; !ok {
		return false, nil
	}

	external, ok := node[externalKey].(bool)
	if !ok {
		return false, errors.Errorf("%s was of unexpected type %T", externalKey, node[externalKey])
	}

	return external, nil
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

	components, ok := spec[componentsKey].(map[interface{}]interface{})
	if !ok {
		return errors.Errorf("components was of unexpected type %T", spec[componentsKey])
	}

	parameters, ok := components[parametersKey].(map[interface{}]interface{})
	if !ok {
		return errors.Errorf("parameters was of unexpected type %T", spec[parametersKey])
	}

	// Build up a record of all query parameters, and then specifically the ones marked external.
	// The distinction is important because non-query parameters when they appear in an endpoint
	// must be left alone
	allQueryParams := map[string]struct{}{}
	externalQueryParams := map[string]struct{}{}
	for parameterKey, parameter := range parameters {
		parameterMap, ok := parameter.(map[interface{}]interface{})
		if !ok {
			return errors.Errorf("components.parameters.%s was of unexpected type %T", parameterKey, parameter)
		}

		in, ok := parameterMap[inKey].(string)
		if !ok {
			return errors.Errorf("components.parameters.%s.%s was of unexpected type %T", parameterKey, inKey, parameterMap[inKey])
		}

		if in != queryValue {
			continue
		}

		parameterName := fmt.Sprintf("#/components/parameters/%s", parameterKey)
		allQueryParams[parameterName] = struct{}{}

		external, err := isExternal(parameterMap)
		if err != nil {
			return errors.WithMessagef(err, "components.parameters.%s", parameterKey)
		}

		if !external {
			msg().Msgf("Deleting components.parameters.%s due to %s: absent or false", parameterKey, externalKey)
			delete(parameters, parameterKey)
			continue
		}

		externalQueryParams[parameterName] = struct{}{}
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

			external, err := isExternal(method)
			if err != nil {
				return errors.WithMessagef(err, "paths.%s.%s", pathKey, key)
			}

			if !external {
				msg().Msgf("Deleting path %s.%s due to %s: absent or false", pathKey, key, externalKey)

				deleted++
				delete(pathMap, key)
				continue
			}

			if parameters, ok := method[parametersKey].([]interface{}); ok {
				foundParams := 0
				parametersToDelete := []int{}
				for i, parameter := range parameters {
					foundParams++
					parameterMap, ok := parameter.(map[interface{}]interface{})
					if !ok {
						return errors.Errorf("paths.%s.%s.parameters was of unexpected type %T", pathKey, path, parameter)
					}

					// Is it a reference to a common component?
					ref, ok := parameterMap[refKey].(string)
					if ok {
						if _, ok := allQueryParams[ref]; !ok {
							// This is not a query parameter
							continue
						}

						if _, ok := externalQueryParams[ref]; !ok {
							msg().Msgf("Deleting parameter %s.%s.%s: %s due to %s: absent or false", pathKey, key, refKey, ref, externalKey)
							parametersToDelete = append(parametersToDelete, i)
						}

						continue
					}

					// It is an "inline" parameter

					in, ok := parameterMap[inKey].(string)
					if !ok {
						return errors.Errorf("parameters %s.%s.% was of unexpected type %T", pathKey, key, inKey, parameterMap[inKey])
					}

					if in != queryValue {
						continue
					}

					parameterName := "unnamed"
					if name, ok := parameterMap[nameKey].(string); ok {
						parameterName = name
					}

					external, err := isExternal(parameterMap)
					if err != nil {
						return errors.WithMessagef(err, "paths.%s.%s.%s", pathKey, key, parameterName)
					}

					if !external {
						msg().Msgf("Deleting parameter %s.%s.%s due to %s: absent or false", pathKey, key, parameterName, externalKey)
						parametersToDelete = append(parametersToDelete, i)
					}
				}

				if foundParams == len(parametersToDelete) {
					msg().Msgf("Deleting parameters from %s.%s as it has no remaining parameters", pathKey, key)
					delete(method, parametersKey)
				} else {
					remove := func(slice []interface{}, i int) []interface{} {
						copy(slice[i:], slice[i+1:])
						return slice[:len(slice)-1]
					}
					// Have to remove in reverse order since trampling otherwise occurs
					sort.Sort(sort.Reverse(sort.IntSlice(parametersToDelete)))
					for _, i := range parametersToDelete {
						parameters = remove(parameters, i)
					}
					method[parametersKey] = parameters
				}
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
