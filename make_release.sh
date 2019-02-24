#!/bin/bash

SOURCES="*.go"

mkdir -p ./binary/darwin
mkdir -p ./binary/windows
mkdir -p ./binary/linux

echo "compiling..."
GOOS=darwin  GOARCH=amd64 go build -o ./binary/darwin/etherpad-cli ${SOURCES}
GOOS=linux   GOARCH=amd64 go build -o ./binary/linux/etherpad-cli ${SOURCES}
GOOS=windows GOARCH=amd64 go build -o ./binary/windows/etherpad-cli.exe ${SOURCES}

echo "done!"
