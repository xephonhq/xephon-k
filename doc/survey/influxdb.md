# InfluxDB

- index https://www.influxdata.com/path-1-billion-time-series-influxdb-high-cardinality-indexing-ready-testing/
  - https://github.com/influxdata/influxdb/tree/tsi
  - current implementation https://github.com/influxdata/influxdb/tree/master/tsdb/index/tsi1
  - new proposal https://github.com/influxdata/influxdb/blob/er-tsi-prop/docs/tsm/TSI_CARDINALITY_PROPOSAL.md
    - the original PR of this new proposal? https://github.com/influxdata/influxdb/pull/7186
- store https://docs.influxdata.com/influxdb/v1.2/concepts/storage_engine/

## Path to 1 Billion Time Series: InfluxDB high cardinality indexing ready for testing

InfluxDB actually looks like two databases in one:

- a time series data store
- an inverted index for the measurement, tag, and field metadata

Old: in memory, build when start
New: TSI file, mmap

- WAL
- LSM
  - In Memory
  - mmap index
  - compact file
- [Robin Hood Hashing](https://github.com/influxdata/influxdb/blob/tsi/pkg/rhh/rhh.go)
  - https://cs.uwaterloo.ca/research/tr/1986/CS-86-14.pdf
- [HyperLogLog++](https://github.com/influxdata/influxdb/blob/tsi/pkg/estimator/hll/hll.go)
  - https://github.com/clarkduvall/hyperloglog

To Solve

similar to [Prometheus's Series Churn](prometheus.md), they also try to solve Kubernetes problem
for https://github.com/kubernetes/heapster I think heapster is using cadvisor, but they don't share storage?

To be Solved

- limit
- having all these series hot for reads and writes
  - [ ] scale-out clustering?
- queries that hit all series in the DB could potentially blow out the memory usage
  - add guard rails and limits into the language and eventually spill-to-disk query processing

## The InfluxDB Storage Engine and the Time-Structured Merge Tree (TSM)

- [ ] shard in retention policy, auto roll up?
- The timestamps and values are compressed and stored separately using encodings dependent on the data type and its shape
- Timestamp
  - delta encoding
  - run-length encoding
  - simple8b https://github.com/jwilder/encoding/tree/master/simple8b
- Float
  - XOR like Gorilla
- Integer
  - ZigZag Protobuf?
  - simple8b
- Strings
  - Snappy
