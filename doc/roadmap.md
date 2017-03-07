# Roadmap

Version 0.1.0

- single tsdb node
- use advanced Cassandra schema (bucket, partition etc.)
- provide RESTful API for querying and insert data
- support `max`, `sum`, `avg` etc.

Version 0.2.0

- multiple tsdb nodes for sharding
- filter by tag (no full text search on tag value, just exact match)
- provide gRPC server for inserting data

Version 0.3.0

- full text search on tags using Elasticsearch
- support visualization using Grafana

Version 0.4.0

- integrate [tsdb-ql](https://github.com/xephonhq/tsdb-ql) as query language

Version 0.5.0

- benchmark against Heroic, KairosDB, InfluxDB
- use [tsdb-proxy](https://github.com/xephonhq/tsdb-proxy) for support multiple input protocol

Version 0.6.0

- replace Cassandra with custom storage engine, may based on RocksDB
