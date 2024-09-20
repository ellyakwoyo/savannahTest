test:
	go test ./...

build:
	go build main.go

.PHONY: test build