#!/bin/bash

DOCKER_IMAGE_NAME=quay.io/newrelic/synthetics-minion-runner
DOCKER_IMAGE_TAG=2.2.23

help() {
    echo "Usage: go [OPTION]"
    echo "  go help"
    echo "  go setup"
    echo "  go clean"
    echo "  go browser | simple-browser | simple_browser"
    echo "  go script | script-browser | script_browser"
    echo "  go api | script-api | script_api"
    echo "  go all | run | run-all | run_all"
}

setup() {
    docker pull $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
}

clean() {
    rm -rf output-script-api/*
    rm -rf output-script-browser/*
    rm -rf output-simple-browser/*
}

simple-browser() {
    docker run --rm \
        -v `pwd`:/opt/shared:rw \
        --env-file ./env-simple-browser \
        --name runner-container-simple-browser \
        $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
    echo "-------- Exit Code: $? --------"
}

script-browser() {
    docker run --rm \
        -v `pwd`:/opt/shared:rw \
        --env-file ./env-script-browser \
        --name runner-container-script-browser \
        $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
    echo "-------- Exit Code: $? --------"
}

script-api() {
    docker run --rm \
        -v `pwd`:/opt/shared:rw \
        --env-file ./env-script-api \
        --name runner-container-script-api \
        $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
    echo "-------- Exit Code: $? --------"
}

if [[ $1 =~ ^(help|setup|clean)$ ]]; then
    "$@"
elif [[ $1 =~ ^(browser|simple-browser|simple_browser)$ ]]; then
    simple-browser
elif [[ $1 =~ ^(script|script-browser|script_browser)$ ]]; then
    script-browser
elif [[ $1 =~ ^(api|script-api|script_api)$ ]]; then
    script-api
elif [[ $1 =~ ^(all|run|run-all|run_all)$ ]]; then
    simple-browser
    script-browser
    script-api
else
    help
    exit 1
fi