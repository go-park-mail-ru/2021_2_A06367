FROM golang:1.16 AS builder

ENV GO111MODULE=on

WORKDIR /opt/app
COPY . .
RUN go build cmd/main/main.go

CMD ./main
