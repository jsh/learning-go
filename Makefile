SHELL := /bin/bash

export GOPATH := $(PWD)

all: true.in
	bin/newtrue

bin/true: src/newtrue/main.go
	go install newtrue

true.in:
	enable -n true; cp $$(which true) true.in

.PHONY: all
