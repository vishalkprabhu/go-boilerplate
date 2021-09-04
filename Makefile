DOCKER_REGISTRY := tweakhub.jfrog.io/default-docker-virtual
REPOSITORY := vishalkprabhu/repository
VERSION := latest
NAME := go-boilerplate

IMAGE := $(DOCKER_REGISTRY)/$(REPOSITORY)/$(NAME)

test:
	true
build-binary:
	cd src && go mod download && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  -o ../bin/go-boilerplate
run-docker: build-binary
	docker-compose up --build
image:
	docker build -t $(IMAGE):$(VERSION) .

push-image:
	docker push $(IMAGE):$(VERSION)


.PHONY: image push-image test