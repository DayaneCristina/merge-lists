FROM golang

RUN apt-get update && apt-get install -y --no-install-recommends build-essential ca-certificates git make bash

WORKDIR /app

ADD go.mod go.sum ./

ADD . .

CMD go build -o main.go