.PHONY: all build run test fmt vet tidy clean help

BINARY := go-products
BIN_DIR := bin

all: build

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) .

run:
	go run .

test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)

help:
	@echo "Targets:"
	@echo "  build  - собрать бинарник в $(BIN_DIR)/$(BINARY)"
	@echo "  run    - запустить программу (go run)"
	@echo "  test   - go test"
	@echo "  fmt    - go fmt"
	@echo "  vet    - go vet"
	@echo "  tidy   - go mod tidy"
	@echo "  clean  - удалить $(BIN_DIR)"
	@echo "  help   - эта справка"
