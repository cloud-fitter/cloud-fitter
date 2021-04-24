#!/bin/bash

set -e
echo "" > coverage.txt

for d in $(go list ./internal/...); do
	echo $d
#    部分sdk内部有并发问题，无法通过检查
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done