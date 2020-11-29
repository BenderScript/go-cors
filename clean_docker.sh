#!/bin/bash

# One liner to stop / remove all of Docker containers:

# docker stop $(docker ps -a -q)
# docker rm $(docker ps -a -q)

CONTAINER_NAME=go-cors

# docker ps -a | awk '{ print $1,$2 }' | grep ${CONTAINER_NAME} | awk '{print $1 }' | xargs -I {} docker rm {}

docker stop ${CONTAINER_NAME} 2> /dev/null || true
docker rm ${CONTAINER_NAME} 2> /dev/null || true
docker rmi -f ${CONTAINER_NAME} 2> /dev/null || true

# https://stackoverflow.com/questions/15678796/suppress-shell-script-error-messages