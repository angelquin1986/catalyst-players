.PHONY: help build run test clean docker-build docker-run docker-stop docker-logs dev-setup

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker Compose"
	@echo "  docker-stop  - Stop Docker containers"
	@echo "  docker-logs  - Show Docker logs"
	@echo "  dev-setup    - Setup development environment"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"

# Build the application
build:
	@echo "Building catalyst-players..."
	go build -o bin/catalyst-players cmd/main.go

# Run the application locally
run:
	@echo "Running catalyst-players..."
	go run cmd/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t catalyst-players .

# Run with Docker Compose (production)
docker-run:
	@echo "Starting production environment..."
	docker compose up -d

# Run with Docker Compose (development)
docker-run-dev:
	@echo "Starting development environment..."
	docker compose -f docker-compose.dev.yml up -d

# Stop Docker containers
docker-stop:
	@echo "Stopping Docker containers..."
	docker compose down
	docker compose -f docker-compose.dev.yml down

# Show Docker logs
docker-logs:
	@echo "Showing Docker logs..."
	docker compose logs -f

# Show Docker logs (development)
docker-logs-dev:
	@echo "Showing development Docker logs..."
	docker compose -f docker-compose.dev.yml logs -f

# Setup development environment
dev-setup:
	@echo "Setting up development environment..."
	@if [ ! -f .env ]; then \
		cp env.example .env; \
		echo "Created .env file from template"; \
	else \
		echo ".env file already exists"; \
	fi
	go mod download
	go mod tidy

# Setup production environment
prod-setup:
	@echo "Setting up production environment..."
	@if [ ! -f .env ]; then \
		cp env.example .env; \
		echo "Created .env file from template"; \
	else \
		echo ".env file already exists"; \
	fi

# Setup development environment with dev-specific config
dev-setup-env:
	@echo "Setting up development environment with dev-specific config..."
	@if [ ! -f .env ]; then \
		cp env.dev.example .env; \
		echo "Created .env file from dev template"; \
	else \
		echo ".env file already exists - manually update DB_HOST=mysql-dev and DB_NAME=catalyst_soccer_dev if needed"; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Generate API documentation
docs:
	@echo "Generating API documentation..."
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g cmd/main.go; \
	else \
		echo "swag not found. Install it with: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

# Database operations
db-migrate:
	@echo "Running database migrations..."
	go run cmd/main.go

db-reset:
	@echo "Resetting database..."
	docker compose down -v
	docker compose up -d

# Health check
health:
	@echo "Checking application health..."
	@curl -f http://localhost:8080/api/v1/health || echo "Application is not running"

# Performance test
bench:
	@echo "Running benchmarks..."
	go test -bench=. ./...

# Coverage
coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html" 