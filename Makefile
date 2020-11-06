#!/bin/sh

PROJECTNAME=$(shell basename "$(PWD)")

BUILD_DIR := ./bin

clean: ## Clean build files
	@echo "  >  Cleaning build cache"
	@go clean
	@-rm -rf $(BUILD_DIR)

build: ## Build the executable
	@echo "  >  Building ${PROJECTNAME}"
	@mkdir -p $(BUILD_DIR)
	@if [ ! -e go.mod ]; then go mod init ${PROJECTNAME}; fi
	@go build -o bin/pizzabot cmd/pizzabot.go

run: ## Run the executable
	./bin/pizzabot $(ARGS)
