# Survey

Survey of existing Time series databases design and implementation details, schema, data structure,
storage engine, on disk format, index format etc. 
Contents are merged with [Xephon-S](https://github.com/xephonhq/xephon-s/tree/master/doc/survey), 
[tsdb-proxy-java](https://github.com/xephonhq/tsdb-proxy-java/blob/master/doc/survey)


Cassandra

- [Cassandra](cassandra.md)
- [ScyllaDB](scylladb.md)

[Time Series Databases using Cassandra as Backend](https://github.com/xephonhq/awesome-time-series-database#cassandra)

- [KairosDB](kairosdb.md)
- [OpenTSDB](opentsdb.md)
- [Heroic](heroic.md)
- [Newts](newts.md)
- [Khronus](khronus.md)
- [Blueflood](blueflood.md)

## TODO


## Usage

- `docker run --name tsdb-cassandra -p 9042:9042 -d cassandra:3.9`
  - `docker stop tsdb-cassandra`
  - `docker start tsdb-cassandra`
- `docker exec -it tsdb-cassandra bash`
- `docker run --name tsdb-scylla -p 9042:9042 -d scylladb/scylla:1.6.0`

## Stuff

To be organized

- RRD can't back fill http://graphite.readthedocs.io/en/latest/whisper.html#differences-between-whisper-and-rrd

## Done

- [x] Thrift, CQL and the underlying storage (I think I got a bit confused when trying to use
  CQL to understand KairosDB's schema design), see [Cassandra](cassandra/README.md)