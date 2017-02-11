# Cassandra

This folder contains reading notes for several (old but useful) official documentations
and blogs from datastax. It's hard to merge them to one file

Datastax

- CQL
  - http://www.datastax.com/dev/blog/cql3-for-cassandra-experts
    - [x] [Schema in Cassandra 1.1](schema-in-cassandra-1-1.md) [link](http://www.datastax.com/dev/blog/schema-in-cassandra-1-1)
    - http://www.datastax.com/dev/blog/whats-new-in-cql-3-0
  - [Does CQL support dynamic columns / wide rows (2013)](http://www.datastax.com/dev/blog/does-cql-support-dynamic-columns-wide-rows)
    - > Thus, the way to model dynamic cells in CQL is with a compound primary key. For the gory details on things like CompositeType, see my previous post.
- Data model
  - [x] [Data Model](1.0-about-data-model.md) [link](http://docs.datastax.com/en/archived/cassandra/1.0/docs/ddl/index.html)
- Storage engine
  - [x] [Storage engine 3.0 (2015)](3.0-storage-engine.md) [link](http://www.datastax.com/2015/12/storage-engine-30)
  - [ ] http://stackoverflow.com/questions/34570367/cassandra-3-0-updated-sstable-format
  - [ ] http://thelastpickle.com/blog/2016/03/04/introductiont-to-the-apache-cassandra-3-storage-engine.html
- Time series
  - [ ] [Basic time series with Cassandra](http://www.rubyscale.com/post/143067470585/basic-time-series-with-cassandra)
    - server-1-load-20110306, put date in the row key, and only have time in column
    - small physical row, no larger than 10MB
  - [ ] [Advanced time series with Cassandra](http://www.datastax.com/dev/blog/advanced-time-series-with-cassandra)
    - use materialized view, for simplying storing integer value, this is useless, for more than of field, this is useful
    - meta row for timeline starting points, for query that only have end without a start
    - different split factor to share load, the split factor can be different at different time, need meta row
    - variable bucket size to avoid big row and sparse row, need meta row as well
    - use separated process to handle meta row to avoid race, then coordinate might be needed
      - [ ] NOTE: I think the security policy of Fedora makes it hard to play with docker ....
  - [ ] [Metric Collection and Storage with Cassandra](http://www.datastax.com/dev/blog/metric-collection-and-storage-with-cassandra)
    - Datastax has OpsCenter for monitoring Cassandra, though it is not open source, and no longer support OSS version of C*

stackoverflow

- [ ] http://stackoverflow.com/questions/24949676/difference-between-partition-key-composite-key-and-clustering-key-in-cassandra
- [ ] http://stackoverflow.com/questions/15857779/commitlog-and-sstables-in-cassandra-database
  - it tells sstable structure
