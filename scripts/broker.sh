#!/bin/bash

if [ "$(basename $(realpath .))" != "go-mqtt" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

COMMAND="$1"
SUBCOMMAND="$2"

export MOSQUITTO_IMAGE="eclipse-mosquitto:2.0.14"
export NETWORK_NAME="go-mqtt_network"

function clean(){
    docker rmi -f ${MOSQUITTO_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function solo(){
    local cmd=$1
    case $cmd in
        "start")
            docker compose -f ./deployment/solo/docker-compose.yaml up
            ;;
        "stop")
            docker compose -f ./deployment/solo/docker-compose.yaml down
            ;;
        *)
            echo "Usage: $0 solo [start|stop]"
            ;;
    esac
}

case $COMMAND in
    "clean")
        clean
        ;;
    "solo")
        solo $SUBCOMMAND
        ;;
    *)
        echo "Usage: $0 [solo]"
        ;;
esac
