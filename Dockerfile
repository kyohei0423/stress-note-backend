FROM golang:1.15.0-alpine3.12 as build

ENV APP="/go/src/github.com/kyohei0423/stress-note-backend"
WORKDIR ${APP}
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    git && \
    go get github.com/pilu/fresh
