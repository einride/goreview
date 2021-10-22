SHELL := /bin/bash

.PHONY: all
all: \
	commitlint \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy \
	yaml-format \
	markdown-format \
	git-verify-nodiff

include tools/commitlint/rules.mk
include tools/golangci-lint/rules.mk
include tools/semantic-release/rules.mk
include tools/prettier/rules.mk
include tools/git-no-diff/rules.mk

.PHONY: clean
clean:
	$(info [$@] removing build files...)
	@rm -rf build

.PHONY: go-review
go-review:
	$(info [$@] reviewing Go code...)
	@go run . ./...

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@mkdir -p build/coverage
	@go test -short -race -coverprofile=build/coverage/$@.txt -covermode=atomic ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v
