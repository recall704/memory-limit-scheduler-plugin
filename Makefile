
NAME=memory-limit-scheduler-plugin
PROJECT_NAME=github.com/recall704/memory-limit-scheduler-plugin
BUILD_IMG=golang:1.20.11
BASE_IMG=debian:stretch-slim

IMAGE_NAME=win7/${NAME}
VERSION=0.0.1


.PHONY: vet
vet:
	@echo ">> vetting code"
	go vet -mod=vendor ./...

build:
	- env GOOS=linux GOARCH=amd64 go build -o bin/${NAME} main.go

docker-build:
	- sed "s|PROJECT_NAME|${PROJECT_NAME}|g" Dockerfile.tpl > Dockerfile
	- sed -i "s|NAME|${NAME}|g" Dockerfile
	- sed -i "s|BUILD_IMG|${BUILD_IMG}|g" Dockerfile
	- sed -i "s|BASE_IMG|${BASE_IMG}|g" Dockerfile
	- cat Dockerfile
	- docker build \
		-t $(IMAGE_NAME):$(VERSION) .

docker-it:
	- docker run -it --rm \
		-v $(PWD):/go/src/$(PROJECT_NAME) \
		-w /go/src/$(PROJECT_NAME) \
		$(BUILD_IMG) \
		bash
