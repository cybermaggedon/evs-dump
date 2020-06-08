
VERSION=$(shell git describe | sed 's/^v//')
REPO=cybermaggedon/evs-dump
DOCKER=docker
GO=GOPATH=$$(pwd)/go go

all: evs-dump build

evs-dump: evs-dump.go go.mod go.sum
	${GO} build evs-dump.go

build: evs-dump
	${DOCKER} build -t ${REPO}:${VERSION} -f Dockerfile .

push:
	${DOCKER} push ${REPO}:${VERSION}

