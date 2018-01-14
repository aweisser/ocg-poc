linux:
	GOOS=linux GOARCH=amd64 go build -o 'ocg-rest-linux-amd64' ./cmd/ocg-rest-server/main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o 'ocg-rest-windows-amd64' ./cmd/ocg-rest-server/main.go

release: windows linux

.PHONY release windows linux: