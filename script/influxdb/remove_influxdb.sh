#!/usr/bin/env bash

docker stop tsdb-influxdb
docker rm tsdb-influxdb
sudo rm -rf /home/at15/workspace/ssd/lib/influxdb