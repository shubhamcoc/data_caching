# Dockerfile for Cache Service
FROM ubuntu:18.04
LABEL description="cache image"

ARG GO_VERSION
RUN apt-get update && \
    apt-get install -y wget git && \
    wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOPATH "/EM/go/src"
ENV PATH $PATH:/usr/local/go/bin:${GOPATH}/bin

RUN go get github.com/eclipse/paho.mqtt.golang
RUN go mod init github.com/my/repo && \
    go get github.com/go-redis/redis/v8 && \
    go get -u github.com/go-redis/cache/v8    

RUN apt-get update && \
    apt-get -y install redis-server

WORKDIR /EM/go/src/
COPY . /usr/local/go/src/cache/
RUN go build -o /usr/local/go/bin/cache /usr/local/go/src/cache/main.go

ENTRYPOINT ["/usr/local/go/bin/cache"]

