FROM golang:1.16-alpine

WORKDIR /go/src/cli
COPY . .

RUN apk add --no-cache build-base
RUN apk add --no-cache jq
RUN apk add --no-cache bash
RUN apk add --no-cache curl

RUN go build -o todo-cli

# Nothing to do... Define `command` in docker-compose.yml
