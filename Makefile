.PHONY: default docs

default: docs synctera-go-client.tar.gz

docs: doc/internal-api.html doc/external-api.html

doc/%-api.html: spec/%-api-merged-bundled.yml
	redoc-cli bundle $<
	mv redoc-static.html $@

merge-external-config = merge-external-apis.json
$(shell jq -r .output $(merge-external-config)): $(shell jq -r .inputs[].inputFile $(merge-external-config))
	openapi-merge-cli --config $(merge-external-config)

merge-internal-config = merge-internal-apis.json
$(shell jq -r .output $(merge-internal-config)): $(shell jq -r .inputs[].inputFile $(merge-internal-config))
	openapi-merge-cli --config $(merge-internal-config)

spec/%-api-bundled.yml: spec/%/api.yml spec/%/*.yml spec/*/*.yml spec/*/*/*.yml
	openapi bundle $< --ext yml --output $@

package-name = synctera

external-%-client: spec/external-api-merged-bundled.yml
	./generate-client.sh external $*

internal-%-client: spec/internal-api-merged-bundled.yml
	./generate-client.sh internal $*

synctera-%-client.tar.gz: external-%-client
	tar -C client/external/ --transform "s|^$*|synctera|" -czf $@ --exclude-from client/external/$*/.tar.ignore $*/

docker-%:
	docker run --user $(shell id -u):$(shell id -g) \
		--mount type=bind,source=$(shell go env GOMODCACHE),destination=/go/pkg/mod,readonly \
		--mount type=bind,source=$(CURDIR),destination=/openapi \
		--workdir /openapi \
		--pull always \
		registry.gitlab.com/synctera/tools/build-tools \
		make $*