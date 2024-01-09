FROM golang

RUN apt-get update && apt-get install -y --no-install-recommends build-essential ca-certificates git make bash

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

ADD go.mod go.sum ./

ADD . .

CMD go build -o main.go