
SHELL := /bin/bash

GOCMD=go
GORUN=$(GOCMD) run

SERVICE=article

init:
	$(GOCMD) mod init $(SERVICE)

tidy:
	$(GOCMD) mod tidy

run:
	$(GORUN) main.go	

migration-up:
	$(GORUN) main.go --migrate-up

migration-down:
	$(GORUN) main.go --migrate-down