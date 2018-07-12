FROM alpine:3.8
RUN apk update && apk add --no-cache git make
ENV GOPATH="/go"
WORKDIR /go/src/github.com/ntrrg/ntgo
COPY . ./
VOLUME /go/src/github.com/ntrrg/ntgo

