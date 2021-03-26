# Dockerfile for Backend service
FROM ubuntu:18.04
LABEL description="Backend image"

ARG GO_VERSION
RUN apt-get update && \
    apt-get install -y wget git && \
    wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOPATH "/EM/go:/EM/go/src"
ENV PATH $PATH:/usr/local/go/bin:${GOPATH}/bin

WORKDIR /EM/go/src/
COPY . ./backend

ENV GO111MODULE "on"
RUN cd backend && go mod init && go mod tidy && cd ..
RUN cd backend && go build -o /EM/go/bin/backend ./cmd/backend-server/main.go

ENTRYPOINT ["/EM/go/bin/backend"]

