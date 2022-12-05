#!/usr/bin/env bash

echo -e "\n>>_StyleChecking_<<"
go fmt ./...
if [[ $? -gt 0 ]]; then exit 1; fi

echo -e "\n>>_Linting_<<"
golangci-lint run ./...
if [[ $? -gt 0 ]]; then exit 1; fi

echo -e "\n>>_UnitTests_<<"
go test -vet=off -count=1 -race -coverprofile=coverage.test ./...
if [[ $? -gt 0 ]]; then exit 1; fi

echo -e "\n>>_GenerateCoveragePage_<<"
go tool cover -html=coverage.test -o ./index.html
if [[ $? -gt 0 ]]; then exit 1; fi

rm ./coverage.test
echo -e "\n>>_ChecksEnded_<<\n"
