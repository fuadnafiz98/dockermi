DEFAULT_GOAL := help

APP := dockermi
BIN_DIR := bin
CMD := ./cmd
PKGS := ./...
ARCH ?= amd64
GO ?= go

help:
	@echo "tidy          Sync module files"
	@echo "deps          Download dependencies"
	@echo "fmt           Format code"
	@echo "vet           Static analysis"
	@echo "lint          Run golangci-lint if available"
	@echo "test          Run tests"
	@echo "cover         Run tests with coverage"
	@echo "build         Build local binary"
	@echo "run           Run app"
	@echo "install       Install to GOBIN"
	@echo "clean         Remove build artifacts"
	@echo "build-linux   Cross-compile for linux ($(ARCH))"
	@echo "build-darwin  Cross-compile for darwin ($(ARCH))"
	@echo "build-windows Cross-compile for windows ($(ARCH))"

tidy:
	$(GO) mod tidy

deps:
	$(GO) mod download

fmt:
	$(GO) fmt $(PKGS)

vet:
	$(GO) vet $(PKGS)

lint:
	@command -v golangci-lint >/dev/null 2>&1 && golangci-lint run || echo "golangci-lint not installed"

test:
	$(GO) test -v $(PKGS)

cover:
	$(GO) test -coverprofile=coverage.out $(PKGS)
	$(GO) tool cover -func=coverage.out

build:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/$(APP) $(CMD)

run:
	$(GO) run $(CMD)

install:
	$(GO) install $(CMD)

clean:
	rm -rf $(BIN_DIR) coverage.out coverage.html

build-linux:
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=$(ARCH) $(GO) build -o $(BIN_DIR)/$(APP)-linux-$(ARCH) $(CMD)

build-darwin:
	mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=$(ARCH) $(GO) build -o $(BIN_DIR)/$(APP)-darwin-$(ARCH) $(CMD)

build-windows:
	mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=$(ARCH) $(GO) build -o $(BIN_DIR)/$(APP)-windows-$(ARCH).exe $(CMD)

.PHONY: help tidy deps fmt vet lint test cover build run install clean build-linux build-darwin build-windows
