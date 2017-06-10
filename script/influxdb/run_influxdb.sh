#!/usr/bin/env bash

docker start tsdb-influxdb
sleep 5
curl -XPOST 'http://localhost:8086/query?u=myusername&p=mypassword' --data-urlencode 'q=CREATE DATABASE "sb"'