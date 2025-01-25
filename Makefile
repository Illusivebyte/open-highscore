# Makefile for the open-highscore project

# Variables
APP_NAME = open-highscore
GO_CMD = go
GO_BUILD = $(GO_CMD) build
GO_TEST = $(GO_CMD) test
GO_RUN = $(GO_CMD) run
GOFMT = $(GO_CMD) fmt

# Default target (run the application)
.PHONY: default
default: run

# Build the application
.PHONY: build
build:
	$(GO_BUILD) -o $(APP_NAME) ./cmd/main.go

# Run the application
.PHONY: run
run:
	$(GO_RUN) ./cmd/main.go

# Run tests
.PHONY: test
test:
	$(GO_TEST) ./...

# Format the code
.PHONY: fmt
fmt:
	$(GOFMT) ./...

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(APP_NAME)

# Lint the code (if you have a linter like golangci-lint)
.PHONY: lint
lint:
	golangci-lint run

# Generate the documentation (using go doc or a similar tool)
.PHONY: docs
docs:
	# Add custom doc generation if desired
	go doc ./...

# Install dependencies (if you're using a dependency management tool)
.PHONY: install-deps
install-deps:
	$(GO_CMD) mod tidy
