#!/bin/bash

GOOS=linux
GOARCH=amd64
export GOOS GOARCH
go build -v -o redis-test-$GOOS-$GOARCH

# rsync -avz  redis-test-$GOOS-$GOARCH cdvm:/tmp/

