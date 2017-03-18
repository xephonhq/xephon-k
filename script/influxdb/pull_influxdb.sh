#!/usr/bin/env bash

docker run --name tsdb-influxdb -p 8083:8083 -p 8086:8086 -d influxdb:1.2.2