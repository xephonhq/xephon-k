#!/usr/bin/env bash

# TODO: merge the scripts
# TODO: mount volume
docker run --name tsdb-cassandra -p 9042:9042 -d cassandra:3.10