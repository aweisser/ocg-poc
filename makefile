BINARY_NAME_LINUX=ocg-rest-linux-amd64

default: clean deps generate test lint release

test:
	go test -v -cover ./...

lint:
	gometalinter --config=.gometalinter.json `find . -not -path "./vendor" -not -path "./vendor/*" -not -path "./io/rest/*" -type d`
	lll --goonly --maxlength 100 --exclude "//go:generate" . --skiplist app design vendor .git
clean:
	go clean
	rm -f $(BINARY_NAME_LINUX)

linux:
	GOOS=linux GOARCH=amd64 go build -o '$(BINARY_NAME_LINUX)' ./cmd/ocg-rest-server/main.go

release: linux

deps:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/goadesign/goa/...
	dep ensure
	gometalinter --install

generate:
	go generate ./...

.PHONY test lint clean linux deps generate release: