DOCKER_REGISTRY := tweakhub.jfrog.io/default-docker-virtual
REPOSITORY := vishalkprabhu/repository
VERSION := latest
NAME := go-boilerplate

IMAGE := $(DOCKER_REGISTRY)/$(REPOSITORY)/$(NAME)

test:
	true

image:
	docker build -t $(IMAGE):$(VERSION) .

push-image:
	docker push $(IMAGE):$(VERSION)


.PHONY: image push-image test