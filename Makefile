GOCMD=go
GOBUILD=$(GOCMD) build

.PHONY: all
all: build

build:
	$(GOBUILD) -o main ./cmd/main

run: build
	./main
