[![Build Status](https://travis-ci.org/aweisser/ocg-poc.svg?branch=master)](https://travis-ci.org/aweisser/ocg-poc)
[![Go Report Card](https://goreportcard.com/badge/github.com/aweisser/ocg-poc)](https://goreportcard.com/report/github.com/aweisser/ocg-poc)

# ocg-poc

Welcome to OCG.

## Setup

### Install dependencies for development
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/matryer/moq
	go get -u golang.org/x/tools/cmd/goimports
	
### Install gometalinter linters
	gometalinter --install

### Install dependencies for production (e.g. those used as imports)
	dep ensure

## First Steps

### Run tests
	go test -v -cover ./...

### Exeute gometalinter for all dirs
	gometalinter --config=.gometalinter.json `find . -not -path "./vendor" -not -path "./vendor/*" -type d`
	
### Execute main command
	go run cmd/ocg-rest-server/main.go