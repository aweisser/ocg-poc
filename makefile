BINARY_NAME_LINUX=ocg-rest-linux-amd64

default: clean deps test lint release

test:
	go test -v -cover ./...

lint:
	gometalinter --config=.gometalinter.json `find . -not -path "./vendor" -not -path "./vendor/*" -not -path "./io/rest/*" -type d`

clean:
	go clean
	rm -f $(BINARY_NAME_LINUX)
	rm -f $(BINARY_NAME_WINDOWS)

linux:
	GOOS=linux GOARCH=amd64 go build -o '$(BINARY_NAME_LINUX)' ./cmd/ocg-rest-server/main.go

release: linux

deps:
	go get github.com/golang/dep/cmd/dep
	go get github.com/alecthomas/gometalinter
	dep ensure
	gometalinter --install


.PHONY test lint clean windows linux release: