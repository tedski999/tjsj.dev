.PHONY: all
all: check build

.PHONY: build
build: bin/sitegen bin/webserver
	@echo -e "\n# Build complete."

.PHONY: check
check:
	@echo -e "\n# Running unit tests..."
	go test -cover ./pkg/sitegen ./pkg/webserver
	@echo -e "\n# Tests complete."

.PHONY: clean
clean:
	@echo -e "\n# Removing build..."
	rm -rf bin

bin/%: cmd/% $(shell find pkg -type f) Makefile
	@echo -e "\n# Building $@..."
	go build  -o $@ ./$<
