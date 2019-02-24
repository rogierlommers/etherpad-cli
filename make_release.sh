#!/bin/bash

SOURCES="*.go"

echo "compiling..."
GOOS=darwin  GOARCH=amd64 go build -o ./binary/etherpad-cli-macos ${SOURCES}
GOOS=linux   GOARCH=amd64 go build -o ./binary/etherpad-cli-linux ${SOURCES}
GOOS=windows GOARCH=amd64 go build -o ./binary/etherpad-cli.exe ${SOURCES}
echo "done!"
