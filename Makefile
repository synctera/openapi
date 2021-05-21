.PHONY: specs

specs: spec/external-api-merged-bundled.yml spec/internal-api-merged-bundled.yml

merge-external-config = merge-external-apis.json
$(shell jq -r .output $(merge-external-config)): $(shell jq -r .inputs[].inputFile $(merge-external-config))
	openapi-merge-cli --config $(merge-external-config)

merge-internal-config = merge-internal-apis.json
$(shell jq -r .output $(merge-internal-config)): $(shell jq -r .inputs[].inputFile $(merge-internal-config))
	openapi-merge-cli --config $(merge-internal-config)

spec/%/api-bundled.yml: spec/%/api.yml
	openapi bundle $^ --ext yml --output $@
