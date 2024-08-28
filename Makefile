.DEFAULT_GOAL := help

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: buf-generate
buf-generate: ## Generate source files from protobuf sources
	@buf generate

.PHONY: build-protos
build-protos: buf-generate ## Build/generate source files from protobuf sources

.PHONY: build
build: build-protos ## Build
	go build -o ./server .

.PHONY: test
test: ## Run tests
	go test -race -v -count=1 -timeout=10m ./...