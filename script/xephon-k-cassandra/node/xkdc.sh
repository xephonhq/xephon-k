#!/usr/bin/env bash

echo "waiting for cassandra to start"
# FIXME： timeout is not included
wait-for-it -t 60 xkbxephonkcassandracassandra:9042
#sleep 10
echo "cassandra started"
# FIXME: copy config file and use new command line flags
xkd schema --cassandra-host xkbxephonkcassandracassandra
xkd schema --cassandra-host xkbxephonkcassandracassandra
xkd schema --cassandra-host xkbxephonkcassandracassandra
xkd -b c --cassandra-host xkbxephonkcassandracassandra
