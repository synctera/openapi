#!/usr/bin/env bash
set -e

tar --transform "s|^|synctera/|" -cf synctera.tar --exclude publish.sh --exclude *_test.go *
gzip synctera.tar
