# Use Alpine base images to reduce image size
FROM golang:1.16-alpine

WORKDIR /go/rest
COPY . .

RUN go build -o todo-server

CMD ./todo-server
