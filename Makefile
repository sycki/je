PWD := $(shell pwd)
SYSTEM := $(shell uname -s)-amd64
VERSION := v$(shell cat version)
OPTS := -ldflags "-X main.version=$(VERSION)"

default: all

all: clean bin tar

clean: 
	@rm -rf _output/*

bin:
	@echo "build je $(VERSION)"
	@GOOS=linux go build $(OPTS) -o _output/je-linux-amd64 ./cmd/je
	@GOOS=darwin go build $(OPTS) -o _output/je-darwin-amd64 ./cmd/je
	@echo "successful binary to _output/"

tar:
	@cd _output && tar -zcf je-linux-amd64.tar.gz je-linux-amd64
	@cd _output && tar -zcf je-darwin-amd64.tar.gz je-darwin-amd64
	@echo "successful tarball to _output/"

