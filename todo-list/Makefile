APP_NAME := todo-list
BUILD_DIR := bin

.PHONY: help run build test fmt vet clean

help:
	@echo "Available targets:"
	@echo "  make run    - Run the app"
	@echo "  make build  - Build binary into ./$(BUILD_DIR)"
	@echo "  make test   - Run tests"
	@echo "  make fmt    - Format Go code"
	@echo "  make vet    - Run go vet"
	@echo "  make clean  - Remove build artifacts"

run:
	go run .

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) .

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -rf $(BUILD_DIR)
