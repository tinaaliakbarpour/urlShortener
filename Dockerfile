FROM golang:1.16-alpine

RUN apk add build-base

RUN apk add --no-cache bash

RUN mkdir -p /go/src/application 

WORKDIR /go/src/application 

ADD . .

RUN go get ./... && go mod vendor && go mod verify


CMD ["./run.sh"]