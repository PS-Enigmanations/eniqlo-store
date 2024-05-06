-include .env

ADDR := localhost:8000
PROJECTNAME := $(shell basename "$(PWD)")
DATABASE_URL := "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?${DB_PARAMS}"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## run golang-migrate up
migrateup:
	migrate -database $(DATABASE_URL) -path db/migrations up

## run golang-migrate down
migratedown:
	migrate -database $(DATABASE_URL) -path db/migrations down