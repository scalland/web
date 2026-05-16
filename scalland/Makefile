APP_NAME := 

.PHONY: all build clean test run serve help db-migrate

# Default target
all: build

# Build the executable
build:
	@echo "Fetching dependencies..."
	@go mod tidy
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) main.go

# Start the server locally
serve:
	@echo "Starting $(APP_NAME) server..."
	@go run main.go serve

# Run database migrations manually
db-migrate:
	@echo "Running migrations..."
	@go run main.go migrate

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
