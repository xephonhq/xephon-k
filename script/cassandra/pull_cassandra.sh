#!/usr/bin/env bash

# TODO: merge the scripts
# TODO: use not hard coded path
docker run --name tsdb-cassandra -v /home/at15/workspace/ssd/lib/cassandra:/var/lib/cassandra -p 9042:9042 -d cassandra:3.10