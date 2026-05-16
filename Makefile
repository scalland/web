APP_NAME := scalland

VERSION := $(shell cat VERSION 2>/dev/null || echo "dev")
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown")
GIT_STATE := $(shell if test -z "$$(git status --porcelain 2>/dev/null)"; then echo "clean"; else echo "dirty"; fi)
GIT_SUMMARY := $(shell git describe --tags --always --dirty 2>/dev/null || echo "none")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -ldflags="-X 'main.Version=$(VERSION)' -X 'main.GitCommit=$(GIT_COMMIT)' -X 'main.GitBranch=$(GIT_BRANCH)' -X 'main.GitState=$(GIT_STATE)' -X 'main.GitSummary=$(GIT_SUMMARY)' -X 'main.BuildDate=$(BUILD_DATE)'"

.PHONY: all build clean test run serve help db-migrate

# Default target
all: build

# Build the executable
build:
	@echo "Fetching dependencies..."
	@go mod tidy
	@echo "Building $(APP_NAME)..."
	@go build $(LDFLAGS) -o $(APP_NAME) main.go

# Start the server locally
serve:
	@echo "Starting $(APP_NAME) server..."
	@go run $(LDFLAGS) main.go serve

# Run database migrations manually
db-migrate:
	@echo "Running migrations..."
	@go run $(LDFLAGS) main.go migrate

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)

# Test project logic
test:
	@echo "Running tests..."
	@go test -v ./...

help:
	@echo "Available commands:"
	@echo "  build       - Fetch dependencies and compile the binary"
	@echo "  serve       - Run the native HTTP web server application"
	@echo "  db-migrate  - Run the database schema migrations"
	@echo "  test        - Run Go tests"
	@echo "  clean       - Clean outputs"
