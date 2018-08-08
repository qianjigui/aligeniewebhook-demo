PWD:= $(shell pwd)

GO:=export GOPATH=$(PWD) && go 

all: compile
	bin/main | tee m.log

compile:
	$(GO) install ./...

test:
	$(GO) test ./...
