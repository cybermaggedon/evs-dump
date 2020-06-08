
VERSION=$(shell git describe | sed 's/^v//')
REPO=cybermaggedon/evs-dump
DOCKER=docker
GO=GOPATH=$$(pwd)/go go

all: evs-dump build

deps: go go/src/github.com/cybermaggedon/evs-golang-api

evs-dump: evs-dump.go
	${GO} build evs-dump.go

build: evs-dump
	${DOCKER} build -t ${REPO}:${VERSION} -f Dockerfile .

push:
	${DOCKER} push ${REPO}:${VERSION}

