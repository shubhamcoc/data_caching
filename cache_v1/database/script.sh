#!/bin/bash

mkdir -p /tmp/influxdb/log

influxd -config /etc/influxdb/influxdb.conf &> /tmp/influxdb/log/influxd.log &
