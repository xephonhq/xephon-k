#!/usr/bin/env bash

# https://docs.timescale.com/other-sample-datasets
# psql -U postgres 
# create database devices_small;
# \c devices_small;
# CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
# psql -U postgres -d devices_small < devices.sql
# psql -U postgres -d devices_small -c "\COPY readings FROM devices_small_readings.csv CSV"
# psql -U postgres -d devices_small -c "\COPY device_info FROM devices_small_device_info.csv CSV"
docker run -d --name tsdb-timescaledb -p 5432:5432 \
  -v /home/at15/workspace/data:/opt/xephonk \
  -e PGDATA=/var/lib/postgresql/data/timescaledb \
  timescale/timescaledb postgres \
  -cshared_preload_libraries=timescaledb