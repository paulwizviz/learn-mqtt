#!/bin/bash

export GRPC_IMAGE=learn-mqtt/grpc

COMMAND="$1"

message="$0 [build | clean | proto ]"

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "build")
        docker-compose -f ./build/grpc/builder.yaml build
        ;;
    "proto")
        docker run -v "$PWD/protos":"/opt/protos" -v "$PWD/internal/":"/opt/internal/" -w /opt -it --rm ${GRPC_IMAGE} /bin/bash -c /opt/protos/scripts/gen.sh
        go mod tidy
        ;;
    "clean")
        docker rmi -f ${GRPC_IMAGE}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    *)
        echo $message
        ;;
esac