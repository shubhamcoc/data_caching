# Dockerfile for Cache Service
FROM ubuntu:18.04
LABEL description="cache image"

ARG GO_VERSION
RUN apt-get update && \
    apt-get install -y wget git && \
    wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOPATH "/EM/go:/EM/go/src"
ENV PATH $PATH:/usr/local/go/bin:${GOPATH}/bin

RUN apt-get update && \
    apt-get -y install redis-server

WORKDIR /EM/go/src/
COPY . /EM/go/src/cache/
RUN cd cache && go mod init && go mod tidy
RUN cd cache && go build -o /usr/local/go/bin/cache ./cmd/cache-server/main.go

ENTRYPOINT ["/usr/local/go/bin/cache"]

