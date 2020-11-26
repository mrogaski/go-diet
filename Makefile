GOCMD=go

.PHONY: test clean

default: all

all: fmt vet test

fmt:
		$(GOCMD) fmt ./...

vet:
		$(GOCMD) vet ./...

test:
		$(GOCMD) test -v ./...

clean:
		$(GOCMD) clean
