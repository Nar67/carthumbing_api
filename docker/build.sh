#!/usr/bin/env bash

docker build -t hackupc2017w/carthumbing_api -f Dockerfile . \
&& ./docker/docker-compose build

if test $? -eq 0; then
    echo "Build successful"
    docker images | grep ^carthumbing_api
else
    echo "ERROR IN BUILD"
fi