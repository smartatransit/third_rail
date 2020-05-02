#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

export CGO_ENABLED=0
export GO111MODULE=on
export GOFLAGS="-mod=vendor"

TARGETS=$(for d in "$@"; do echo ./$d/...; done)

echo "Running tests:"
t="/tmp/go-cover.$$.tmp"
go test -v -installsuffix "static" ${TARGETS} -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
echo

