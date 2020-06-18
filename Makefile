
# Create version tag from git tag
VERSION=$(shell git describe | sed 's/^v//')
REPO=cybermaggedon/evs-dump
DOCKER=docker
GO=GOPATH=$$(pwd)/go go

all: evs-dump build

evs-dump: evs-dump.go go.mod go.sum
	${GO} build evs-dump.go

build: evs-dump
	${DOCKER} build -t ${REPO}:${VERSION} -f Dockerfile .
	${DOCKER} tag ${REPO}:${VERSION} ${REPO}:latest

push:
	${DOCKER} push ${REPO}:${VERSION}
	${DOCKER} push ${REPO}:latest


