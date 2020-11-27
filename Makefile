GOCMD=go

.PHONY: fmt vet test clean

default: all

all: fmt vet test

fmt:
		$(GOCMD) fmt ./...

vet:
		$(GOCMD) vet ./...

test:
		$(GOCMD) test ./... -v -coverprofile=coverage.txt -covermode=atomic

clean:
		$(GOCMD) clean
