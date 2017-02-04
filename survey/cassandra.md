# Cassandra

## Usage

For Fedora `sudo systemctl start docker.service` might be needed

- `docker run --name tsdb-cassandra -p 9042:9042 -d cassandra:3.9`
  - `docker stop tsdb-cassandra`
  - `docker start tsdb-cassandra`
- `docker exec -it tsdb-cassandra bash`

## TODO:

- [ ] keys http://stackoverflow.com/questions/24949676/difference-between-partition-key-composite-key-and-clustering-key-in-cassandra

## Compact storage

- https://docs.datastax.com/en/cql/3.3/cql/cql_reference/cqlCreateTable.html#refClstrOrdr__cql-compact-storage
- **For Cassandra 3.0 and later, the storage engine is much more efficient at storing data, and compact storage is not necessary.**
  - [ ] TODO: http://www.datastax.com/2015/12/storage-engine-30

## Wide row

- [ ] TODO: underlying storage
- [ ] TODO: in old thrift way
- [ ] TODO: in CQL

[Does CQL support dynamic columns / wide rows (2013)](http://www.datastax.com/dev/blog/does-cql-support-dynamic-columns-wide-rows)

> Thus, the way to model dynamic cells in CQL is with a compound primary key. For the gory details on things like CompositeType, see my previous post.


- http://www.datastax.com/dev/blog/cql3-for-cassandra-experts
  - http://www.datastax.com/dev/blog/schema-in-cassandra-1-1
  - http://www.datastax.com/dev/blog/whats-new-in-cql-3-0
- http://stackoverflow.com/questions/24949676/difference-between-partition-key-composite-key-and-clustering-key-in-cassandra


http://www.datastax.com/dev/blog/schema-in-cassandra-1-1

- Cassandra used to be kind of schema less
  - > Using UUIDs as a surrogate key is common in Cassandra, so that you don’t need to worry about sequence or autoincrement synchronization across multiple machines
- [ ] TODO: OpenTSDB has said Cassandra does not handle UUID generation
  - > Additionally, because Cassandra doesn't support atomic mutations, locking for atomic increments is implemented by writing a special lock column and checking the timestamps to see who won with retries if acquisition failed. That means UID assignments will be messy
