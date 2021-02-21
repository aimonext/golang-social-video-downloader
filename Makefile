# Makefile for Go Social Video Downloader

# Variables
APP_NAME=go-social-video-downloader
BUILD_DIR=bin
VERSION=1.0.0

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/main.go
	@echo "Build complete! Executable created in $(BUILD_DIR)/"

# Build a statically linked binary
static-build:
	@echo "Building static $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build -ldflags="-s -w" -a -installsuffix cgo -o $(BUILD_DIR)/$(APP_NAME) ./cmd/main.go
	@echo "Static build complete! Executable created in $(BUILD_DIR)/"

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	@go run cmd/main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@echo "Dependencies installed!"

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  run      - Run the application"
	@echo "  clean    - Clean build artifacts"
	@echo "  deps     - Install dependencies"
	@echo "  help     - Show this help message"

.PHONY: build run clean deps help
