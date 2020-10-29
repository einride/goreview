.PHONY: all
all: \
	commitlint \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy

include tools/golangci-lint/rules.mk
include tools/commitlint/rules.mk
include tools/semantic-release/rules.mk

.PHONY: go-review
go-review:
	$(info [$@] reviewing Go code...)
	@go run ./cmd/goreview/main.go ./...

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@go test -race -cover ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v
