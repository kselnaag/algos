#!/usr/bin/env bash

echo -e "\n>>_Formating_<<"
go fmt ./...

echo -e "\n>>_Linting_<<"
golangci-lint run ./...

echo -e "\n>>_Testing_<<"
go test -vet=off -count=1 -race -coverprofile=coverage.test ./...
go tool cover -html=coverage.test -o ./coverage.html
rm ./coverage.test
