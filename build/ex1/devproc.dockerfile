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

ARG GO_VER

FROM golang:${GO_VER} as builder 

ARG APP_NAME

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.sum ./go.sum
COPY ./go.mod ./go.mod

RUN go mod download && \
    env CGO_ENABLED=0 env GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/${APP_NAME} ./cmd/ex1/devproc

FROM ubuntu:18.04

ARG APP_NAME

COPY --from=builder /opt/build/${APP_NAME} /${APP_NAME}

CMD /${APP_NAME}