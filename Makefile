DOCKER_REGISTRY := tweakhub.jfrog.io/default-docker-virtual
REPOSITORY := vishalkprabhu/repository
VERSION := latest
NAME := go-boilerplate

IMAGE := $(DOCKER_REGISTRY)/$(REPOSITORY)/$(NAME)

WORKSPACE := `pwd`


.PHONY: image push-image test

test:
	true

build-binary:
	cd src && env GOCACHE=$(WORKSPACE)/tmp GOPATH=$(WORKSPACE)/go/pkg  go build -o ../bin/$(NAME)
	cd ..

run-docker:
	cd src && env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o ../bin/$(NAME)
	cd ..
	docker-compose up --build

image:
	docker build -t $(IMAGE):$(VERSION) .

push-image:
	docker push $(IMAGE):$(VERSION)

create-db-local:
	docker exec gp-mysql mysql -uroot -ppassword -e "CREATE DATABASE IF NOT EXISTS mydb"

migrate-up-local:
	migrate -source file://migrations -database "mysql://root:password@tcp(localhost:3306)/mydb" up

migrate-down-local:
	migrate -source file://migrations -database "mysql://root:password@tcp(localhost:3306)/mydb" down