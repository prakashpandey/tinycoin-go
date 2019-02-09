GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
BINARY_NAME=go-chain-$(GOOS)-$(GOARCH)

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)

.PHONY: build
build:
	go build -o $(BINARY_NAME)

.PHONY: run
run: clean fmt build
	./$(BINARY_NAME)

.PHONY: fmt
fmt:
	go fmt ./...
