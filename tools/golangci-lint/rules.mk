golangci_lint_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
golangci_lint := $(golangci_lint_cwd)/bin/golangci-lint

$(golangci_lint): $(golangci_lint_cwd)/go.mod
	cd $(golangci_lint_cwd) && go build -o $@ github.com/golangci/golangci-lint/cmd/golangci-lint
