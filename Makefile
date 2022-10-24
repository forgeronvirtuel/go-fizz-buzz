.PHONY: build

dep:
	go mod download
	go mod verify

build:
	mkdir -p ./build
	go build -o build/fizzbuzz cmd/server.go

run:
	go run cmd/server.go