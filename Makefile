.PHONY: all
all: \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy

include tools/golangci-lint/rules.mk

.PHONY: go-lint
go-lint: $(golangci_lint)
	$(golangci_lint) run

.PHONY: go-review
go-review:
	go run ./cmd/goreview/main.go ./...

.PHONY: go-test
go-test:
	go test -race -cover ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v
