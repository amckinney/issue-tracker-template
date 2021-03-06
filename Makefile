.DEFAULT_GOAL := run

.PHONY: run
run: install
	go run ./cmd/istkr/main.go

.PHONY: build
build: install
	rm -f bin/istkr
	go build ./... -o bin/istkr

.PHONY: install
install:
	go mod download

.PHONY: upgrade
upgrade:
	go get -u ./...

.PHONY: format
format: gofmt

.PHONY: gofmt
gofmt:
	gofmt -s -w $(shell find . -name '*.go')

.PHONY: clean
clean:
	git clean -xdf
