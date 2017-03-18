#!/usr/bin/env bash

echo "influx"
echo "create database sb"
echo "show field keys;"
docker exec -it tsdb-influxdb bash