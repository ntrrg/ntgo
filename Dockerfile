FROM golang:1.10.3-alpine
RUN apk update && apk add --no-cache git make
RUN go get -u gopkg.in/alecthomas/gometalinter.v2 && gometalinter.v2 --install
WORKDIR /go/src/github.com/ntrrg/ntgo
VOLUME /go/src/github.com/ntrrg/ntgo

