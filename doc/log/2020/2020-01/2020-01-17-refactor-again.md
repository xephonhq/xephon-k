# 2020-01-17 Refactor again

It's almost (exactly) 1 year after [previous effort for refactoring](../../2019/2019-01/2019-01-25-0.0.3-init-refactor.md) (well rewrite).
I have to say although I didn't write any code, I did write a detailed plan in [#69](https://github.com/xephonhq/xephon-k/issues/69).

## Design

I'd start with the simplest data model and single node

Data model

- tagged time series like Prometheus and InfluxDB

Storage engine

- assuming tag index can fit into memory (obviously it may not be the case)
- use simple columnar store, i.e. just blocks.
- support WAL but can be disabled (mainly used to see the effect of having both WAL and data on single disk)
- support compaction on single node

Compression

- relies on libtsdb-go

Protocol

Each protocol should have a prepared statement style variant (because is might be much faster? though gzip entire payload could beat it?)

- http json (kairosdb like)
- tcp, text and binary
- grpc
- prometheus

Query

- filter by tag and returns as it is, i.e. no aggregation, no group by. Mainly used for verify correctness

## Dependencies

- compression in libtsdb
  - need to move existing simple compression code to libtsdb, and implement gorilla tsz
- error and logging from gommon
- tracing?
  - I want to have a simple in process tracing, it may not be that feature rich [tokio-rs/tracing](https://github.com/tokio-rs/tracing) might be a good reference.
- end to end benchmark from xephon-b