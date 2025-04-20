.PHONY: build run clean install test

# Default binary name
BINARY_NAME=sims4-mod-manager

# Build the application
build:
	@echo "Building ${BINARY_NAME}..."
	@go build -o ${BINARY_NAME} -v

# Run the application with UI
run: build
	@echo "Running ${BINARY_NAME}..."
	@./${BINARY_NAME} ui

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f ${BINARY_NAME}
	@go clean

# Install the application to GOPATH/bin
install: build
	@echo "Installing ${BINARY_NAME}..."
	@go install

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-windows-amd64.exe
	@GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME}-darwin-amd64
	@GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux-amd64

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Show help
help:
	@echo "Sims 4 Mod Manager - Make commands"
	@echo "make build        - Build the application"
	@echo "make run          - Build and run the application with UI"
	@echo "make clean        - Clean build artifacts"
	@echo "make install      - Install the application to GOPATH/bin"
	@echo "make build-all    - Build for multiple platforms"
	@echo "make test         - Run tests"
	@echo "make help         - Show this help"

# Default target
default: build