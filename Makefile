all: help

help:
	@echo "available commands:"
	@echo "  - make build           build the binary"

build:
	CGO_ENABLED=0 go build -o fhttp .
