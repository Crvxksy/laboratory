FROM golang:1.17

MAINTAINER Ivory Air

ENV GOPATH /go
ENV GO111MODULE on

COPY air /go/bin/

ENTRYPOINT ["/go/bin/air"]
