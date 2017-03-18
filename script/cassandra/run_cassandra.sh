#!/usr/bin/env bash

# TODO: if docker start fail, pull the image and name it tsdb-cassandra
# FIXME: Fedora need sudo for docker while many other Linux distros not
# TODO: in Fedora, I need sudo for client application to access docker mapped port
# NOTE: may need to switch to vagrant ....
docker start tsdb-cassandra

# TODO: create keyspace, because gocql can not connect without setting keyspace