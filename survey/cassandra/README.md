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
  - [ ] [Storage engine 3.0 (2015)](3.0-storage-engine.md) [link](http://www.datastax.com/2015/12/storage-engine-30)

stackoverflow

- http://stackoverflow.com/questions/24949676/difference-between-partition-key-composite-key-and-clustering-key-in-cassandra
- http://stackoverflow.com/questions/15857779/commitlog-and-sstables-in-cassandra-database
  - it tells sstable structure
