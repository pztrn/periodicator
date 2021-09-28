#!/usr/bin/env bash

# This script is using for building release binaries.

OSES=("darwin" "linux" "windows")
ARCHES=("amd64" "arm64")

GO=$(which go)
if [ "${GO}" == "" ]; then
	echo "Golang is not installed, cannot continue"
fi

if [ ! -d "release" ]; then
	mkdir -p release
fi

for OS in "${OSES[@]}"; do
	for ARCH in "${ARCHES[@]}"; do
		echo "Building for ${OS} ${ARCH}..."
		CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -o release/periodicator-${OS}-${ARCH} .
	done
done

for file in release/*windows*; do
	mv ${file} ${file}.exe
done

for file in release/*; do
	zip -m ${file}.zip ${file}
done
