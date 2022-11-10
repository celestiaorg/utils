SHELL=/usr/bin/env bash
PROJECTNAME=$(shell basename "$(PWD)")
LDFLAGS="-X 'main.buildTime=$(shell date)' -X 'main.lastCommit=$(shell git rev-parse HEAD)' -X 'main.semanticVersion=$(shell git describe --tags --dirty=-dev)'"
ifeq (${PREFIX},)
	PREFIX := /usr/local
endif

# Set the default target for `make`
.DEFAULT_GOAL := help

## benchmark: Running all benchmarks
benchmark:
	@echo "--> Running benchmarks"
	@go test -tags='testing' -run="none" -bench=. -benchtime=100x -benchmem ./...
.PHONY: benchmark

## cover: generate to code coverage report.
cover:
	@echo "--> Generating Code Coverage"
	@go install github.com/ory/go-acc@latest
	@go-acc -o coverage.txt `go list ./... | grep -v nodebuilder/tests` -- -v -tags='testing'
.PHONY: cover

## deps: install dependencies.
deps:
	@echo "--> Installing Dependencies"
	@go mod download
.PHONY: deps

## fmt: Formats only *.go (excluding *.pb.go *pb_test.go). Runs `gofmt & goimports` internally.
fmt:
	@find . -name '*.go' -type f -not -path "*.git*" -not -name '*.pb.go' -not -name '*pb_test.go' | xargs gofmt -w -s
	@find . -name '*.go' -type f -not -path "*.git*"  -not -name '*.pb.go' -not -name '*pb_test.go' | xargs goimports -w -local github.com/celestiaorg
	@go mod tidy -compat=1.17
	@markdownlint --fix --quiet --config .markdownlint.yaml .
.PHONY: fmt

## help: Get more info on make commands.
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help

## hooks: Install git-hooks from .githooks directory.
hooks:
	@echo "--> Installing git hooks"
	@git config core.hooksPath .githooks
.PHONY: hooks

## lint: Linting *.go files using golangci-lint. Look for .golangci.yml for the list of linters.
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@markdownlint --config .markdownlint.yaml '**/*.md'
.PHONY: lint

## test: Running all tests
test:
	@echo "--> Running all tests with data race detector"
	@go test -v -tags='testing' -race ./...
.PHONY: test
