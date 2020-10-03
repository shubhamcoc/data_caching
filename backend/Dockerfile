# Dockerfile for Backend service
FROM ubuntu:18.04
LABEL description="Backend image"

ARG GO_VERSION
RUN apt-get update && \
    apt-get install -y wget git && \
    wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOPATH "/EM/go"
ENV PATH $PATH:/usr/local/go/bin:${GOPATH}/bin

RUN go get github.com/eclipse/paho.mqtt.golang
WORKDIR /EM/go/src/
COPY . ./backend/
RUN go build -o /EM/go/bin/backend /EM/go/src/backend/backend.go

ENTRYPOINT ["backend"]

