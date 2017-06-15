#!/usr/bin/env bash

docker stop tsdb-cassandra
docker rm tsdb-cassandra
sudo rm -rf /home/at15/workspace/ssd/lib/cassandra