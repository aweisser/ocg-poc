BINARY_NAME=ocg-rest

default: deps test release
test:
	go test -v -cover ./...

linux:
	GOOS=linux GOARCH=amd64 go build -o '$(BINARY_NAME)-linux-amd64' ./cmd/ocg-rest-server/main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o '$(BINARY_NAME)-windows-amd64' ./cmd/ocg-rest-server/main.go

release: windows linux

deps:
	go get github.com/golang/dep/cmd/dep
	dep ensure
.PHONY test release windows linux: