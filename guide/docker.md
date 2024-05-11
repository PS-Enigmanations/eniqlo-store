# Guide

### List of all container images

```sh
docker image ls
```

### Remove container image

```sh
docker image rm -f <IMAGE_ID> #or

# conflict: image has dependent child images
docker rmi -f <REPOSITORY_NAME>:<TAG_NAME>
```

### Inspect on running

```sh
docker inspect <CONTAINER_NAME>
```

## Create database

```sh
docker exec -it enigmanations_postgres_container psql -h localhost -p 5430 -U postgres -c 'create database "eniqlo-store"'
```

## Migrate local database to docker

```sh
cat db/init.sql | docker exec -i enigmanations_postgres_container psql -h localhost -p 5432 -U postgres -d eniqlo-store
```

## Run and connect to the database

This allows you to run multiple commands or perform interactive tasks within the container.

```sh
docker exec -it enigmanations_postgres_container bash
```

### Build and running postgres

```sh
docker-compose up --build postgres
```

### Build and running api

```sh
docker-compose up --build api
```

### Build and running all

```sh
docker-compose up --build
```

### Check current active container

```sh
docker ps -a
```
