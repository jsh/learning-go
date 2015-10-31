SHELL := /bin/bash

export GOPATH := $(PWD)

all: true.in bin/newtrue bmk
	bin/newtrue

bmk:
	[ -f bmk ] || touch bmk

bin/newtrue: src/newtrue/main.go
	go install newtrue

true.in:
	enable -n true; cp $$(which true) true.in

clean:
	rm *.OUT

.PHONY: all
