.PHONY: default docs

default: docs external-go-client

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

spec/%-api-bundled.yml: spec/%/api.yml spec/%/*.yml spec/%/*/*.yml spec/common/*.yml
	openapi bundle $< --ext yml --output $@

package-name = synctera

external-%-client: spec/external-api-merged-bundled.yml client/%.config.json
	openapi-generator-cli generate --strict-spec true --generator-name $* \
		--input-spec spec/external-api-merged-bundled.yml --output client/external/$*/ \
		--package-name $(package-name) --config client/$*.config.json

internal-%-client: spec/internal-api-merged-bundled.yml client/%.config.json
	openapi-generator-cli generate --strict-spec true --generator-name $* \
		--input-spec spec/internal-api-merged-bundled.yml --output client/internal/$*/ \
		--package-name $(package-name) --config client/$*.config.json
