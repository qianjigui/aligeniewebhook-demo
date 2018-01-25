PWD:= $(shell pwd)

GO:=export GOPATH=$(PWD) && go 

all: compile
	bin/main

compile:
	$(GO) install ./...

test:
	$(GO) test ./...
