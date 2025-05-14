# Run the application
run:
	go run cmd/server/main.go

# Build the application
build:
	go build -o bin/app cmd/server/main.go

# Install dependencies
install:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Run tests
test:
	go test ./...

# Lint the code (if using golangci-lint)
lint:
	golangci-lint run

# Run all checks before committing
check: fmt lint test
