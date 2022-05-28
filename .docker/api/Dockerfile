FROM golang:1.17.5-alpine3.15

ENV CGO_ENABLED=0

RUN apk update && apk upgrade && \
    apk add --update figlet && \
    go get github.com/cespare/reflex && \
    go get github.com/swaggo/swag/cmd/swag && \
    go get github.com/golang/mock/mockgen && \
    go get github.com/gin-gonic/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/gin-contrib/cors && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/fgrosse/goldi
