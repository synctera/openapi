# OpenAPI

This repository is for creating documentation and clients from OpenAPI 3 specifications.

## Makefile targets

`make docs` creates HTML documentation

`make internal-$lang-client` creates an internal client library for the language `$lang`. For example, `make internal-typescript-axios-client`.

`make external-$lang-client` creates an external client library for the language `$lang`. For example, `make external-go-client`.

`make synctera-$lang-client.tar.gz` creates an external client library package for the language `$lang`

`make docker-$target` runs the target `$target` in a docker container, so you don't have to install any tools locally. For example, `make docker-synctera-go-client.tar.gz`.

## Layout

`spec/` contains the source OpenAPI 3 specification files (`spec/*/api.yml`), their dependent specs and their generated bundled versions (`spec/*-api-bundled.yml*`)

The specifications combine into internal and external packages according to `merge-external-apis.json` and `merge-internal-apis.json`.

`doc/` contains generated HTML documentation

`client/internal/$lang` contains the generated internal client for the language `$lang`

`client/external/$lang` contains the generated external client for the language `$lang`

`client/$lang.config.json` specifies language-specific options for the language `$lang`. See the [OpenAPI generator options](https://openapi-generator.tech/docs/generators) for details.
