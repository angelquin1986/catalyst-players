#!/bin/bash

# Catalyst Players Development Script
# This script helps with development setup and common tasks

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to check prerequisites
check_prerequisites() {
    print_status "Checking prerequisites..."
    
    if ! command_exists go; then
        print_error "Go is not installed. Please install Go 1.21 or later."
        exit 1
    fi
    
    if ! command_exists docker; then
        print_warning "Docker is not installed. Some features may not work."
    fi
    
    if ! command_exists docker-compose; then
        print_warning "Docker Compose is not installed. Some features may not work."
    fi
    
    print_success "Prerequisites check completed"
}

# Function to setup development environment
setup_dev() {
    print_status "Setting up development environment..."
    
    # Create .env file if it doesn't exist
    if [ ! -f .env ]; then
        cp env.example .env
        print_success "Created .env file from template"
    else
        print_warning ".env file already exists"
    fi
    
    # Install dependencies
    go mod download
    go mod tidy
    
    print_success "Development environment setup completed"
}

# Function to run the application locally
run_local() {
    print_status "Running application locally..."
    
    if [ ! -f .env ]; then
        print_error ".env file not found. Run 'bash scripts/dev.sh setup' first."
        exit 1
    fi
    
    # Export variables from .env file for the go run command
    set -a
    source .env
    set +a
    
    go run cmd/main.go
}

# Function to run with Docker
run_docker() {
    print_status "Running with Docker..."
    
    if ! command_exists docker; then
        print_error "Docker is not installed"
        exit 1
    fi
    
    docker compose up -d
    print_success "Application started with Docker"
    print_status "API available at: http://localhost:8080"
    print_status "Health check: http://localhost:8080/api/v1/health"
}

# Function to run with Docker (development)
run_docker_dev() {
    print_status "Running with Docker (development mode)..."
    
    if ! command_exists docker; then
        print_error "Docker is not installed"
        exit 1
    fi
    
    docker compose -f docker-compose.dev.yml up -d
    print_success "Application started with Docker (development)"
    print_status "API available at: http://localhost:8081"
    print_status "Health check: http://localhost:8081/api/v1/health"
}

# Function to stop Docker containers
stop_docker() {
    print_status "Stopping Docker containers..."
    
    docker compose down
    docker compose -f docker-compose.dev.yml down
    
    print_success "Docker containers stopped"
}

# Function to show logs
show_logs() {
    print_status "Showing Docker logs..."
    
    if [ "$1" = "dev" ]; then
        docker compose -f docker-compose.dev.yml logs -f
    else
        docker compose logs -f
    fi
}

# Function to run tests
run_tests() {
    print_status "Running tests..."
    go test -v ./...
}

# Function to build the application
build_app() {
    print_status "Building application..."
    go build -o bin/catalyst-players cmd/main.go
    print_success "Application built successfully"
}

# Function to clean build artifacts
clean_build() {
    print_status "Cleaning build artifacts..."
    rm -rf bin/
    go clean
    print_success "Build artifacts cleaned"
}

# Function to format code
format_code() {
    print_status "Formatting code..."
    go fmt ./...
    print_success "Code formatted"
}

# Function to run linter
run_linter() {
    print_status "Running linter..."
    
    if command_exists golangci-lint; then
        golangci-lint run
        print_success "Linting completed"
    else
        print_warning "golangci-lint not found. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
    fi
}

# Function to show help
show_help() {
    echo "Catalyst Players Development Script"
    echo ""
    echo "Usage: $0 [COMMAND]"
    echo ""
    echo "Commands:"
    echo "  setup       - Setup development environment"
    echo "  run         - Run application locally"
    echo "  docker      - Run with Docker (production)"
    echo "  docker-dev  - Run with Docker (development)"
    echo "  stop        - Stop Docker containers"
    echo "  logs        - Show Docker logs"
    echo "  logs-dev    - Show Docker logs (development)"
    echo "  test        - Run tests"
    echo "  build       - Build application"
    echo "  clean       - Clean build artifacts"
    echo "  format      - Format code"
    echo "  lint        - Run linter"
    echo "  check       - Check prerequisites"
    echo "  help        - Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 setup        # Setup development environment"
    echo "  $0 run          # Run locally"
    echo "  $0 docker-dev   # Run with Docker (development)"
    echo "  $0 test         # Run tests"
}

# Main script logic
case "${1:-help}" in
    setup)
        check_prerequisites
        setup_dev
        ;;
    run)
        run_local
        ;;
    docker)
        run_docker
        ;;
    docker-dev)
        run_docker_dev
        ;;
    stop)
        stop_docker
        ;;
    logs)
        show_logs
        ;;
    logs-dev)
        show_logs dev
        ;;
    test)
        run_tests
        ;;
    build)
        build_app
        ;;
    clean)
        clean_build
        ;;
    format)
        format_code
        ;;
    lint)
        run_linter
        ;;
    check)
        check_prerequisites
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        print_error "Unknown command: $1"
        echo ""
        show_help
        exit 1
        ;;
esac 