BINARY ?= calculator
OUTDIR := bin

.PHONY: all build run clean tidy fmt vet test help

all: build

build:
	@mkdir -p $(OUTDIR)
	go build -o $(OUTDIR)/$(BINARY) .

run: build
	$(OUTDIR)/$(BINARY)

clean:
	rm -rf $(OUTDIR)

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

help:
	@echo "Targets: build (default), run, clean, tidy, fmt, vet, test"
