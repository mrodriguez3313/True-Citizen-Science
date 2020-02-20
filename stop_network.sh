#!/bin/bash

dockerContainers=$(docker ps -a | awk '$2~/ipfs/ {print $1}')
if [ "$dockerContainers" != "" ]; then
    echo "Deleting existing docker containers ..."
    docker rm -f $dockerContainers
fi
