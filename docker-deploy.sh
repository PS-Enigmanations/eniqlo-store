#!/bin/bash

DOCKER_USERNAME="natserract"
REPOSITORY_NAME="enigmanations-inventory"
IMAGE_TAG="$(git log --format="%h" -n 1)"
DOCKERFILE="Dockerfile"

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin && \
    docker build \
        -t ${DOCKER_USERNAME}/${REPOSITORY_NAME}:${IMAGE_TAG} \
        -t ${DOCKER_USERNAME}/${REPOSITORY_NAME}:latest -f ${DOCKERFILE} . && \
    docker push ${DOCKER_USERNAME}/${REPOSITORY_NAME}
