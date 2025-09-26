# TikTok Business API Go SDK Makefile

.PHONY: help build test test-verbose test-coverage clean lint fmt vet mod-tidy mod-download run-example install-tools

# Default target
help: ## Show this help message
	@echo "TikTok Business API Go SDK"
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
# Build targets
build: ## Build the SDK
	@echo "Building Go SDK..."
	go build ./...

# Testing targets
test: ## Run all tests
	@echo "Running tests..."
	go test -v ./...

test-implementation: ## Run implementation quality tests
	@echo "Running implementation quality tests..."
	go test -v ./pkg/client -run TestNoPlaceholdersOrTodos
	go test -v ./pkg/client -run TestAllServicesHaveProperImplementations
	go test -v ./pkg/client -run TestErrorMessagesAreProfessional

test-verbose: ## Run tests with verbose output
	@echo "Running tests with verbose output..."
	go test -v ./...
	@echo "Running tests with coverage..."
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-race: ## Run tests with race detection
	@echo "Running tests with race detection..."
	go test -race ./...

# Code quality targets
lint: install-tools ## Run linter
	@echo "Running linter..."
	golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

# Dependency management
mod-tidy: ## Tidy go modules
	@echo "Tidying go modules..."
	go mod tidy

mod-download: ## Download go modules
	@echo "Downloading go modules..."
	go mod download

mod-verify: ## Verify go modules
	@echo "Verifying go modules..."
	go mod verify

# Development tools
install-tools: ## Install development tools
	@echo "Installing development tools..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@which goimports > /dev/null || (echo "Installing goimports..." && go install golang.org/x/tools/cmd/goimports@latest)

# Example targets
run-example: ## Run the basic usage example
	@echo "Running basic usage example..."
	@echo "Note: Set TIKTOK_ACCESS_TOKEN environment variable for full functionality"
	cd examples/basic_usage && go run main.go

# Cleanup targets
clean: ## Clean build artifacts and coverage files
	@echo "Cleaning up..."
	rm -f coverage.out coverage.html
	go clean ./...

# Documentation targets
docs: ## Generate documentation
	@echo "Generating documentation..."
	godoc -http=:6060 &
	@echo "Documentation server started at http://localhost:6060"
	@echo "Navigate to http://localhost:6060/pkg/github.com/tiktok/tiktok-business-api-sdk/go_sdk/"

# Security targets
security: install-tools ## Run security checks
	@echo "Running security checks..."
	@which gosec > /dev/null || (echo "Installing gosec..." && go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest)
	gosec ./...

# Benchmark targets
bench: ## Run benchmarks
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# All quality checks
check: fmt vet lint test ## Run all quality checks

# CI targets
ci: mod-tidy fmt vet lint test-race test-coverage test-implementation security ## Run all CI checks (Travis CI compatible)

# Travis CI specific target
travis-ci: ci ## Alias for Travis CI compatibility

# Release targets
version: ## Show current version
	@echo "Current version: $(shell git describe --tags --always --dirty)"

tag: ## Create a new tag (usage: make tag VERSION=v1.0.0)
	@if [ -z "$(VERSION)" ]; then echo "Usage: make tag VERSION=v1.0.0"; exit 1; fi
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)

# Development workflow
dev-setup: install-tools mod-download ## Set up development environment
	@echo "Development environment setup complete!"

dev-check: fmt vet lint test ## Quick development checks

# Docker targets (if needed in the future)
docker-build: ## Build Docker image
	@echo "Docker support available in future versions"

docker-test: ## Run tests in Docker
	@echo "Docker support available in future versions"

# Integration test targets
test-integration: ## Run integration tests (requires API credentials)
	@echo "Running integration tests..."
	@echo "Note: Set TIKTOK_ACCESS_TOKEN and other credentials for integration tests"
	go test -tags=integration ./tests/integration/...

# Performance targets
profile-cpu: ## Run CPU profiling
	@echo "Running CPU profiling..."
	go test -cpuprofile=cpu.prof -bench=. ./...
	go tool pprof cpu.prof

profile-mem: ## Run memory profiling
	@echo "Running memory profiling..."
	go test -memprofile=mem.prof -bench=. ./...
	go tool pprof mem.prof

# Generate targets
generate: ## Run go generate
	@echo "Running go generate..."
	go generate ./...

# Dependency analysis
deps: ## Analyze dependencies
	@echo "Analyzing dependencies..."
	go list -m all
	go mod graph

deps-update: ## Update dependencies
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# Project information
info: ## Show project information
	@echo "Project: TikTok Business API Go SDK"
	@echo "Go version: $(shell go version)"
	@echo "Module: $(shell go list -m)"
	@echo "Git commit: $(shell git rev-parse --short HEAD)"
	@echo "Git branch: $(shell git rev-parse --abbrev-ref HEAD)"
	@echo "Build time: $(shell date)"

# Help with common tasks
quick-start: dev-setup ## Quick start for new developers
	@echo ""
	@echo "🚀 Quick Start Guide:"
	@echo "1. Run 'make test' to verify everything works"
	@echo "2. Run 'make run-example' to see the SDK in action"
	@echo "3. Set environment variables for full functionality:"
	@echo "   export TIKTOK_ACCESS_TOKEN='your_token'"
	@echo "   export TIKTOK_CLIENT_ID='your_client_id'"
	@echo "   export TIKTOK_CLIENT_SECRET='your_client_secret'"
	@echo "4. Run 'make dev-check' before committing changes"
	@echo ""
