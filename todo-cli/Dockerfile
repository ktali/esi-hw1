# Use alpine base images to minimize image size
FROM golang:1.16-alpine

WORKDIR /go/src/todo-cli
COPY . .

# Reduce the amount of layers (therefore size) by downloading dependencies with a single RUN directive
RUN apk add --no-cache build-base jq bash curl

RUN go build -o todo-cli

# Nothing to do... Define `command` in docker-compose.yml
