SHELL := /bin/bash

export GOPATH := $(PWD)

all:
	@ echo usage: make '[ badflag | clean | helpflag | noflag | versionflag ]'

badflag: setup
	./true.in --badflag > bmk
	bin/newtrue --badflag

bin/newtrue: src/newtrue/main.go
	go install newtrue

clean:
	rm *.OUT true.out

helpflag: setup
	./true.in --help > bmk
	bin/newtrue --help

noflag: setup
	./true.in > bmk
	bin/newtrue

setup: true.in bin/newtrue

true.in:
	enable -n true; cp $$(which true) true.in

versionflag:  setup
	./true.in --version > bmk
	bin/newtrue --version

.PHONY: all badflag clean help noflag setup version
