FROM golang:1.16-alpine

WORKDIR /go/rest
COPY . .

RUN go build -o todo-server

RUN apk add --no-cache jq

CMD ./todo-server
