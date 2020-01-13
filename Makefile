.PHONY: all
all: \
	mod-tidy \
	go-lint \
	go-review \
	go-test \
	git-verify-nodiff \
	git-verify-submodules

export GO111MODULE = on

.PHONY: build
build:
	git submodule update --init --recursive $@

include build/rules.mk
build/rules.mk: build
	@# Included in submodule: build

# mod-tidy: ensure Go module files are in sync
.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: go-lint
go-lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --config build/.golangci.yml

.PHONY: go-review
go-review:
	go run ./cmd/goreview/main.go ./...

.PHONY: go-test
go-test:
	go test -race -cover ./...

.PHONY: markdown-lint
markdown-lint: $(MARKDOWNLINT)
	$(MARKDOWNLINT) . --ignore vendor --ignore build
