# Survey

This folder contains doc and code snippets from [time series
databases using Cassandra as backend(https://github.com/xephonhq/awesome-time-series-database#cassandra).

C*

- [Cassandra](cassandra.md)
- [ScyllaDB](scylladb.md)

TSDBs

- [KairosDB](kairosdb.md)
- [OpenTSDB](opentsdb.md)
- [Heroic](heroic.md)
- [Newts](newts.md)

## TODO

- Thrift, CQL and the underlying storage (I think I got a bit confused when trying to use
  CQL to understand KairosDB's schema design)

## Usage

- `docker run --name tsdb-cassandra -p 9042:9042 -d cassandra:3.9`
  - `docker stop tsdb-cassandra`
  - `docker start tsdb-cassandra`
- `docker exec -it tsdb-cassandra bash`
- `docker run --name tsdb-scylla -p 9042:9042 -d scylladb/scylla:1.6.0`

## Stuff

To be organized

- RRD can't back fill http://graphite.readthedocs.io/en/latest/whisper.html#differences-between-whisper-and-rrd
