.PHONY: all
all: \
	commitlint \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy

include tools/commitlint/rules.mk
include tools/golangci-lint/rules.mk

.PHONY: go-review
go-review:
	$(info [$@] reviewing Go code...)
	@go run . ./...

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@go test -race -cover ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v
