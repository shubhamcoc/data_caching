# Dockerfile for Database
FROM ubuntu:18.04
LABEL description="InfluxDB image"

ARG GO_VERSION
RUN apt-get update && \
    apt-get install -y wget git && \
    wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOPATH "/EM/go:/EM/go/src"
ENV GO111MODULE on
ENV PATH $PATH:/usr/local/go/bin:${GOPATH}/bin

# Installing mariadb
RUN apt-get install -y mariadb-server
RUN service mysql start

WORKDIR /EM/go/src/
COPY . ./database/
RUN chmod +x ./database/start_db.sh

RUN cd database && go mod init && \
    go mod tidy

RUN cd database && go build -o /EM/go/bin/database ./cmd/database-server/main.go

ENTRYPOINT ["/EM/go/bin/database"]

