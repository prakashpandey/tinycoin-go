GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
BINARY_BASE_NAME=tinycoin-cli
BINARY_NAME=$(BINARY_BASE_NAME)-$(GOOS)-$(GOARCH)

.ONESHELL:
export SHELL:=/bin/bash

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)

.PHONY: build
build:
	go build -o $(BINARY_NAME)

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./target/$(BINARY_BASE_NAME)-$(GOOS)-$(GOARCH)

.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./target/$(BINARY_BASE_NAME)-$(GOOS)-$(GOARCH)


.PHONY: build-windows
build-windows:
	export GOOS=windows \
	export GOARCH=amd64 \
	go build -o ./target/$(BINARY_BASE_NAME)-$(GOOS)-$(GOARCH)


.PHONY: build-all-platforms
build-all-platforms: build-linux build-darwin build-windows
	
.PHONY: run
run: clean fmt build
	./$(BINARY_NAME)

.PHONY: fmt
fmt:
	go fmt ./...
