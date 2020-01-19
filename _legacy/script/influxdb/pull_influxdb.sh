#!/usr/bin/env bash

# TODO: change configuration
# https://docs.influxdata.com/influxdb/v1.2/administration/config/
# default
# cache-max-memory-size = 1G
# cache-snapshot-memory-size = 24MB
# The cache snapshot memory size is the size at which the engine will snapshot the cache and write it to a TSM file, freeing up memor
# INFLUXDB_DATA_CACHE_SNAPSHOT_MEMORY_SIZE
# TODO: use not hard coded path
docker run --name tsdb-influxdb -v /home/at15/workspace/ssd/lib/influxdb:/var/lib/influxdb -p 8083:8083 -p 8086:8086 -d influxdb:1.2.4
sleep 5
curl -XPOST 'http://localhost:8086/query?u=myusername&p=mypassword' --data-urlencode 'q=CREATE DATABASE "xb"'