DOCKER_TAG ?= device-store:latest

.PHONY: build-image run-server

build-image: 
	docker build -t $(DOCKER_TAG) .

run-dev-server: build-image
	docker compose up

test:
	go test -v ./...

init:
	go mod download