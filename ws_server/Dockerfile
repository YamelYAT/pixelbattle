FROM golang:1.15-alpine
RUN apk add build-base
WORKDIR /src

COPY go.mod .
COPY . .

RUN go build -o ./out/ws_server

