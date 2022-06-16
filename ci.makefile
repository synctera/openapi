# NB: this Makefile is intended to be used by CI *after* the ./update script has run successfully

%-api.yml: %-api-merged-bundled.yml
	NODE_PATH=$(shell npm root -g) node ./datafaker.js $< $@

%-api.json: %-api.yml
	openapi bundle $< --ext json --output $@

%-api.html: %-api.yml
	redoc-cli bundle $<
	mv redoc-static.html $@

external-%-code: %.external-api.yml
	./generate-code external go $*

internal-%-code: %.internal-api.yml
	./generate-code internal go $*

go-external.%.tar.gz: external-%-code
	tar -C code/external/go/ --transform "s|^$*|synctera|" -czf $@ --exclude-from code/external/go/$*/.tar.ignore $*/

go-internal.%.tar.gz: internal-%-code
	tar -C code/internal/go/ --transform "s|^$*|synctera|" -czf $@ --exclude-from code/internal/go/$*/.tar.ignore $*/
