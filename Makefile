.PHONY: build run

build:
	@CGO_ENABLE=0 go build -o ./out/server

run: build
	@./out/./server
