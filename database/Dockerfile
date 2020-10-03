# Dockerfile for Database
FROM ubuntu:18.04
LABEL description="InfluxDB image"

ARG GO_VERSION
RUN apt-get update && \
    apt-get install -y wget git && \
    wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOPATH "/EM/go"
ENV PATH $PATH:/usr/local/go/bin:${GOPATH}/bin
ENV INFLUXDB_GO_PATH /EM/go/src/github.com/influxdata/influxdb
RUN mkdir -p ${INFLUXDB_GO_PATH} && \
    git clone https://github.com/influxdata/influxdb ${INFLUXDB_GO_PATH} && \
    cd ${INFLUXDB_GO_PATH} && \
    git checkout -b v1.6.0 tags/v1.6.0

# Installing influxdb
ARG INFLUXDB_VERSION
RUN wget https://dl.influxdata.com/influxdb/releases/influxdb_${INFLUXDB_VERSION}_amd64.deb && \
    dpkg -i influxdb_${INFLUXDB_VERSION}_amd64.deb && \
    rm -rf influxdb_${INFLUXDB_VERSION}_amd64.deb

RUN go get github.com/eclipse/paho.mqtt.golang
WORKDIR /EM/go/src/
COPY . ./database/
RUN go build -o /EM/go/bin/main /EM/go/src/database/main.go

COPY ./influxdb.conf /etc/influxdb/influxdb.conf
ENTRYPOINT ["main"]

