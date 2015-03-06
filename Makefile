PROJECT=companyd-client-go

BUILD_PATH := $(shell pwd)/.gobuild

PROJECT_PATH := "$(BUILD_PATH)/src/github.com/giantswarm"

BIN=$(PROJECT)

.PHONY=clean run-tests run-integration-tests get-deps fmt

GOPATH := $(BUILD_PATH)

SOURCE=$(shell find . -name '*.go')

all: get-deps $(BIN)

clean:
	rm -rf $(BUILD_PATH) $(BIN)

get-deps: .gobuild

.gobuild:
	mkdir -p $(PROJECT_PATH)
	cd "$(PROJECT_PATH)" && ln -s ../../../.. $(PROJECT)

	# Fetch internal or pinned dependencies
	@builder get dep -b status-with-reason git@github.com:giantswarm/api-schema.git $(PROJECT_PATH)/api-schema

	#
	# Fetch public dependencies via `go get`
	GOPATH=$(GOPATH) go get -d -v github.com/giantswarm/$(PROJECT)

$(BIN): $(SOURCE)
	GOPATH=$(GOPATH) go build -o $(BIN)

run-tests:
	GOPATH=$(GOPATH) go test -v -tags "unit" ./...

run-integration-tests:
	# Run compandy locally beforehand on localhost:8080
	GOPATH=$(GOPATH) go test -v -tags "integration" ./...
fmt:
	gofmt -l -w .
