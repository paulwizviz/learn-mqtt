#!/bin/bash

if [ -d /opt/internal/pmodel ]; then
    rm -rf /opt/internal/pmodel
    mkdir /opt/internal/pmodel
else
    mkdir /opt/internal/pmodel
fi
protoc --proto_path=/opt/protos/specs --go_out=/opt/internal/pmodel --go_opt=paths=source_relative /opt/protos/specs/data.proto