.DEFAULT_GOAL := build

.PHONY:fmt vet build run
fmt: 
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build 
ifeq ($(OS),Windows_NT)
RUN_BIN := ./myproject.exe
else
RUN_BIN := ./myproject
endif

run: build
	$(RUN_BIN)