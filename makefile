BINARY_NAME_LINUX=ocg-rest-linux-amd64
BINARY_NAME_WINDOWS=ocg-rest-windows-amd64

default: clean deps test release

test:
	go test -v -cover ./...

clean:
	go clean
	rm -f $(BINARY_NAME_LINUX)
	rm -f $(BINARY_NAME_WINDOWS)

linux:
	GOOS=linux GOARCH=amd64 go build -o '$(BINARY_NAME_LINUX)' ./cmd/ocg-rest-server/main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o '$(BINARY_NAME_WINDOWS)' ./cmd/ocg-rest-server/main.go

release: windows linux

deps:
	go get github.com/golang/dep/cmd/dep
	dep ensure

.PHONY test clean windows linux release: