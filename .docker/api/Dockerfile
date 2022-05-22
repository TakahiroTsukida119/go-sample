FROM golang:1.18-buster

ENV CGO_ENABLED=0
#    GO111MODULE=auto

RUN apt-get update && \
    go install github.com/cespare/reflex@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/golang/mock/mockgen@latest

