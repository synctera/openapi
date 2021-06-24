# NB: this Makefile is intended to be used by CI *after* the ./update script has run successfully

%-api-faked-merged-bundled.yml: %-api-merged-bundled.yml
	NODE_PATH=$(shell npm root -g) node ./datafaker.js $< $@

%-api.html: %-api-faked-merged-bundled.yml
	redoc-cli bundle $<
	mv redoc-static.html $@

external-%-code: external-api-faked-merged-bundled.yml
	./generate-code external $*

internal-%-code: internal-api-faked-merged-bundled.yml
	./generate-code internal $*

%-external.tar.gz: external-%-code
	tar -C code/external/ --transform "s|^$*|synctera|" -czf $@ --exclude-from code/external/$*/.tar.ignore $*/

%-internal.tar.gz: internal-%-code
	tar -C code/internal/ --transform "s|^$*|synctera|" -czf $@ --exclude-from code/internal/$*/.tar.ignore $*/
