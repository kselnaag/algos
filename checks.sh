#!/usr/bin/env bash

echo -e "\n>>_Style-checking_<<"
go fmt ./...

echo -e "\n>>_Linting_<<"
golangci-lint run ./...

echo -e "\n>>_UnitTests_<<"
go test -vet=off -count=1 -race -coverprofile=coverage.test ./...
go tool cover -html=coverage.test -o ./index.html
rm ./coverage.test
