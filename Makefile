PHONY: test fmt build all

all: test build

fmt:
	go fmt ./...

test: fmt
	go test ./...

build:
	go build -o ./dist/main main.go