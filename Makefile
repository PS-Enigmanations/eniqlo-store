-include .env

ADDR := localhost:8080
PROJECTNAME := $(shell basename "$(PWD)")
DATABASE_URL := "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?${DB_PARAMS}"

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

kill:
	lsof -t -i tcp:8080 | xargs kill -9

## dev: run build and up on dev environment.
dev: build up

## prod: run build and up on production environment.
prod: build up-prod

## build: run build on dev environment.
build:
	GOARCH=amd64 GOOS=darwin go build -o main_enigmanations .

## build: run build on production environment.
build-prod:
	GOARCH=amd64 GOOS=linux go build -o main_enigmanations .

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
