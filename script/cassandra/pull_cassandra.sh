#!/usr/bin/env bash

# TODO: merge thes scripts
docker run --name tsdb-cassandra -p 9042:9042 -d cassandra:3.10