-include .env

ADDR := localhost:8080
PROJECTNAME := $(shell basename "$(PWD)")
DATABASE_URL := "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?${DB_PARAMS}"

## Deployment
GIT_HASH ?= $(shell git log --format="%h" -n 1)
DOCKER_USERNAME := "natserract"
APPLICATION_NAME := "enigmanations-inventory"
_BUILD_ARGS_TAG ?= ${GIT_HASH}
_BUILD_ARGS_RELEASE_TAG ?= latest
_BUILD_ARGS_DOCKERFILE ?= Dockerfile

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

setup:
	go install github.com/cosmtrek/air@latest

kill:
	lsof -t -i tcp:8080 | xargs kill -9

## dev: run build and up on dev environment.
dev: build up

## prod: run build and up on production environment.
prod: build-prod up

## build: run build on dev environment.
build:
	GOARCH=amd64 GOOS=darwin go build -o main_enigmanations .

## build: run build on production environment.
build-prod:
	GOARCH=amd64 GOOS=linux go build -o main_enigmanations .

docker-build:
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} -f ${_BUILD_ARGS_DOCKERFILE} .

docker-push:
    docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}

docker-run:
	docker run ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}

## up: run docker-compose up with dev environment.
up:
	JWT_SECRET=a-very-secretive-secret-key ./main_enigmanations

## up: run docker-compose up with production environment.
up-prod:
	ENV=production JWT_SECRET=a-very-secretive-secret-key ./main_enigmanations

## run golang-migrate up
migrateup:
	migrate -database $(DATABASE_URL) -path db/migrations up

## run golang-migrate down
migratedown:
	migrate -database $(DATABASE_URL) -path db/migrations down
