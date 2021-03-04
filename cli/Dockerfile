FROM golang:1.16-alpine

WORKDIR /go/cli
COPY . .

RUN go build -o todo-cli

# to keep the container running I guess
CMD bash
