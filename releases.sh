#!/usr/bin/env bash

buid_and_checksum() {
    env GOOS=$1 GOARCH=amd64 go build -o go-yq-${1}${2}
    sha512sum go-yq-${1}${2} > go-yq-${1}${2}.sha512.txt
}

buid_and_checksum darwin
buid_and_checksum linux
buid_and_checksum windows .exe