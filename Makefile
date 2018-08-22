PWD:= $(shell pwd)

GO:=export GOPATH=$(PWD) && go 

all: compile
	bin/main 2>&1 | tee m.log

compile:
	$(GO) install ./...

test:
	$(GO) test ./...
