# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: docker

docker:
	docker build -t github.com/huyhvq/lingio:$${CI_COMMIT_REF_NAME} .
