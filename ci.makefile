# NB: this Makefile is intended to be used by CI *after* the ./update script has run successfully

%-api.yml: %-api-merged-bundled.yml
	NODE_PATH=$(shell npm root -g) node ./datafaker.js $< $@

%-api.json: %-api.yml
	openapi bundle $< --ext json --output $@

%-api.html: %-api.yml
	redoc-cli bundle $<
	mv redoc-static.html $@

external-%-code: external-api.yml
	./generate-code external $*

internal-%-code: internal-api.yml
	./generate-code internal $*

%-external.tar.gz: external-%-code
	tar -C code/external/ --transform "s|^$*|synctera|" -czf $@ --exclude-from code/external/$*/.tar.ignore $*/

%-internal.tar.gz: internal-%-code
	tar -C code/internal/ --transform "s|^$*|synctera|" -czf $@ --exclude-from code/internal/$*/.tar.ignore $*/
