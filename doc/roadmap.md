# Roadmap

## Current

- support query
- built in graphing web service [#19](https://github.com/xephonhq/xephon-k/issues/19)
  - using echarts
  - pack using gulp
  - (optional) bind assets to the binary (may need to rewrite gorice, I don't like gobinddata)
  - (optional) watch and reload
- support aggregation on query
  - various window size
  - avg
  - sum
  - min
  - max
- support gRPC
- [Survey](survey)
  - [x] [existing TSDBs using C*](https://github.com/xephonhq/awesome-time-series-database#cassandra)
  - [x] categorize schemas
- [Specification](spec-draft.md)
  - [x] naive schema
  - [x] [naive schema's naive implementation](../pkg/bin/xnaive/main.go)
- storing metrics
  - [x] memory without index tag
  - [x] cassandra without index tag
- query metrics as it is
  - [x] without using tag
- query aggregation
- index text without using external search engine
  - i.e. https://github.com/balzaczyy/golucene

## Outline

Version 0.1.0

- single tsdb node
- use advanced Cassandra schema (bucket, partition etc.)
- provide RESTful API for querying and insert data
- support `max`, `sum`, `avg` etc.

Version 0.2.0

- multiple tsdb nodes for sharding
  - split big time range into small time range and aggregate result
  - limit resources for each docker node
- filter by tag (no full text search on tag value, just exact match)
- provide gRPC server for inserting data

Version 0.3.0

- full text search on tags using Elasticsearch
- ~~support visualization using Grafana~~ we will use own graphing system

Version 0.4.0

- integrate [tsdb-ql](https://github.com/xephonhq/tsdb-ql) as query language

Version 0.5.0

- benchmark against Heroic, KairosDB, InfluxDB
- use [tsdb-proxy](https://github.com/xephonhq/tsdb-proxy) for support multiple input protocol

Version 0.6.0

- replace Cassandra with custom storage engine, may based on RocksDB
