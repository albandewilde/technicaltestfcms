.PHONY: build run test

build:
	@CGO_ENABLE=0 go build -o ./out/server

run: build
	@./out/./server

test:
	@go test ./...
