.PHONY: build run test image ctn-run

build:
	@CGO_ENABLE=0 go build -o ./out/server

run: build
	@./out/./server

test:
	@go test ./...

image:
	@docker build . -t technicaltestfcms

ctn-run: image
	@docker container run -p 8254:8254 technicaltestfcms
