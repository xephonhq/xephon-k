#!/usr/bin/env bash

# TODO: merge the scripts
# TODO: mount volume
docker run --name tsdb-scylladb -p 9042:9042 -d scylladb/scylla:1.7.0