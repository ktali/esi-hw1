FROM golang:1.16-alpine

WORKDIR /go/rest
COPY . .

RUN go build -o todo-server

RUN apk add --no-cache jq
RUN apk add --no-cache bash
RUN apk add --no-cache curl

RUN chmod +x functional-test.sh

CMD ./todo-server
