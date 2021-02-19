# Makefile for Go Social Video Downloader

# Variables
APP_NAME=go-social-video-downloader
BUILD_DIR=bin
VERSION=1.0.0

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/main.go
	@echo "Build complete! Executable created in $(BUILD_DIR)/"

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
