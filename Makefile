.DEFAULT_GOAL := help

.PHONY: clean build fmt test

APP_NAME	  = off-api
ROOT_DIR      := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BUILD_FLAGS   ?=
VERSION 	  = $(shell cat version.txt)
BRANCH        = $(shell git rev-parse --abbrev-ref HEAD)
REVISION      = $(shell git describe --tags --always --dirty)
BUILD_DATE    = $(shell date +'%Y.%m.%d-%H:%M:%S')
LDFLAGS       ?= -w -s

BUILD_DIR     = $(ROOT_DIR)/build
BINARY        = $(BUILD_DIR)/$(APP_NAME)

LOCAL_IMAGE   = $(APP_NAME):latest

default: help

.PHONY: help
help:
	@grep -E '^[a-zA-Z%_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build executable
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GO111MODULE=on go build -mod=vendor -o $(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

test: ## Test
	@mkdir -p $(BUILD_DIR)
	GO111MODULE=on go test -mod=vendor -v ./...

.PHONY: check
check:
	staticcheck -go module -f stylish ./...

fmt: ## Go format
	go fmt ./...

clean: ## Clean
	@rm -rf $(BUILD_DIR)

.PHONY: generate
generate: ## Go generate
	GO111MODULE=on go generate ./...

.PHONY: deps
deps: ## Get dependencies
	GO111MODULE=on go get ./...

.PHONY: vendor
vendor: ## Go vendor
	GO111MODULE=on go mod vendor

.PHONY: tidy
tidy: ## Go tidy
	GO111MODULE=on go mod tidy

run:
	GO111MODULE=on go run main.go

