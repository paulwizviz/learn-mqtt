#!/bin/bash

# Copyright 2022 Paul Sitoh
# 
# Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

COMMAND="$1"

export DEVICE_IMAGE_NAME=mqttmock/devicemock:current
export DEVPROC_IMAGE_NAME=mqttmock/devprocmock:current
export DEVOP_IMAGE_NAME=mqttmock/devop:current

export MOSQUITTO_CONTAINER="mosquitto-server"
export MOCK_DEVICE_CONTAINER="mock-device"
export MOCK_DEVPROC_CONTAINER="mock-devproc"
export DEVOP_CONTAINER="devop"

export CLIENT_APP_NAME="mqttclient"

export NETWORK_NAME="mqtt-mock-network"

function build(){
    docker-compose -f ./build/ex1/builder.yaml build device
    docker-compose -f ./build/ex1/builder.yaml build devproc
    docker-compose -f ./build/ex1/builder.yaml build devop
}

function clean(){
    docker rmi -f ${DEVICE_IMAGE_NAME}
    docker rmi -f ${DEVPROC_IMAGE_NAME}
    docker rmi -f ${DEVOP_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function run(){
    docker-compose -f ./deployment/ex1/docker-compose.yml up
}

function devop(){
    docker-compose -f ./deployment/ex1/devop.yml up passwd
}

function stop(){
    docker-compose -f ./deployment/ex1/docker-compose.yml down
}

message="$0 build | clean | devop | run  | stop"

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    "devop")
        devop
        ;;
    "run")
        run
        ;;
    "stop")
        stop
        ;;
    *)
        echo $message
        ;;
esac