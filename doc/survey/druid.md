# Druid

## Take away

- https://dev.twitter.com/streaming/public

## Stories from the Trenches – The Challenges of Building an Analytics Stack

https://www.youtube.com/watch?v=Sz4w75xRrYM

![druid](druid.png)

- Immutable data
- In Memory is Overrated
  - mmap + SSD
  - cost of scaling CPU << cost of adding RAM
  - decompression on the fly (LZF, Snappy, LZ4)
- Low Latency vs High throughput
  - combine batch + streaming
  - immutable made it easy to combine the two ingestion methods
  - Makes for easy backfill and re-processing
  - Historical Node
  - Real-time Node
- Not All Data is Created Equal
  - user really care about recent data
  - user still want to run quarterly report
  - large queries create bottlenecks and resource contention
- Smarter Rebalancing
- Create Data Tiers
- Addressing Multitenancy
  - HyperLogLog sketches
  - Approximate top-k
  - Approximate histograms (monitoring)
- Monitoring
  - Use Druid to monitor Druid
- **Use cases should define engineering**

## Druid: A Real-time Analytical Data Store

- time series data with both numeric and text value
- a set dimension columns
  - KairosDB etc. are not multi dimension, strictly speaking
  - query over any arbitrary combination of dimensions

Features for a dashboard

- query latency
- multi-tenant
- HA
- make business decisions in "real-time"

### Architecture

> The name Druid comes from the Druid class in many role-playing games: it is a
shape-shifter, capable of taking on many different forms to fulfill
various different roles in a group.

- real time node
- historical node
- broker node
- coordinator node

Real-time Node

- use ZK for online state and data range?
- Row store when in JVM heap-based buffer
- persist in-memory indexes to disk periodically/based on threshold of rows
  - [ ] TODO: just indexes? how does Druid know which node to build index
  - column oriented storage format
- [ ] load persisted indexes into off-heap memory
- immutable block **segment** into deep storage (s3, HDFS)
- **consume Kafka**
  - buffer (and crash recovery)
  - single end point for multiple real-time nodes to read events

Historical Nodes

- load, drop, serve immutable segments
- use ZK for online state and data range
- download segment from deep storage, first check local cache
- **group into Tiers**

Broker Nodes

- broker nodes act as query routers to historical and real-time nodes
- merge partial results
- LRU cache

Coordinator Nodes

- use ZK to select leader, other as backup
- MVCC swapping protocol to maintain stable views
- **MySQL**
  - a list of all segments that should be served by historical node
  - [ ] then what is ZK used for?
  - configuration
- **Rules**
  - rules govern how historical segments are loaded and dropped from the cluster
- cost-based optimization for load balancing
- replication
  - used for rolling update
  - the data is actually in deep storage (s3, HDFS)

Storage Format

- **Segement**: a collection of rows of data that span some period of time
  - unit for replication
  - data source identifier
  - time interval
    - [ ] TODO: must be fixed interval, or this is a approximate interval
  - version string for concurrency control
- **druid always requires a timestamp column as a method of simplifying data distribution policies, data retention policies, and first level query pruning**
- time granularity to partition segments is a function of data volume and time range
  - [ ] TODO: how to know the range, let user specify it
  - i.e. data spread a year is partitioned by day, data spread over a day is partitioned by hour
- **Column store**
  - dictionary encoding for string
  - [ ] where is the dictionary stored, in segment?
  - **use generic compression algorithm on top of encoding**
  - LZF
- **Bitmap index**

Storage Engine

- in memory (i.e. JVM heap)
- memory mapped

> main drawback with using the memory-mapped storage engine is when a query requires more segments to be paged into memory
than a given node has capacity for.

- [ ] doesn't JVM heap also have this problem and bigger

Query API

- filter
  - type i.e. selector
  - dimension (column)
  - value
- granularity: i.e. day
- aggregations
  - type i.e. count
  - name i.e. rows

- **don't want to implement join**

### Performance

- TPC-H
- 10TB RAM (wow ..... that's e... well 64 * 160 ... (2^6+2^4) * 10)
