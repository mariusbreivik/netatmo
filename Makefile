VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -s -w \
	-X github.com/mariusbreivik/netatmo/cmd.Version=$(VERSION) \
	-X github.com/mariusbreivik/netatmo/cmd.Commit=$(COMMIT) \
	-X github.com/mariusbreivik/netatmo/cmd.BuildDate=$(BUILD_DATE)

.PHONY: build
build: ## Build the binary with version info
	go build -ldflags "$(LDFLAGS)" -o netatmo .

.PHONY: install
install: ## Install the binary with version info
	go install -ldflags "$(LDFLAGS)" .

.PHONY: test
test: ## Run tests
	go test -v -race ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: lint
lint: ## Run linter
	golangci-lint run

.PHONY: clean
clean: ## Remove build artifacts
	rm -f netatmo coverage.out
	rm -rf dist/

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

