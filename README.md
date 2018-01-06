# ocg-poc

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