# KairosDB

- GitHub: https://github.com/kairosdb/kairosdb

## Protocol

- Thrift
- CQL

## Schema

- [Thrift version](https://github.com/kairosdb/kairosdb/blob/master/src/main/java/org/kairosdb/datastore/cassandra/CassandraDatastore.java#L218)
- [CQL version](https://github.com/kairosdb/kairosdb/blob/feature/cql/src/main/java/org/kairosdb/datastore/cassandra/CassandraDatastore.java)

### Metrics Schema

Keyspace

- NOTE: the replication_factor is 1, and it is a fixed number in `CREATE_KEYSPACE` in `CassandraDatastore.java`

````sql
CREATE KEYSPACE IF NOT EXISTS {{keyspace}}
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
  };
````

Metrics

- NOTE: it seems KairosDB has omitted keyspace when creating tables
- NOTE: column1 is not a typo, it might be the legacy problem of using Thrift protocol

````sql
CREATE TABLE IF NOT EXISTS data_points (
    key blob,
    column1 blob,
    value blob,
    PRIMARY KEY ((key), column1)
  ) WITH COMPACT STORAGE;
````

Row key index

- [ ] TODO: what is it used for, tags?

````sql
CREATE TABLE IF NOT EXISTS row_key_index (
    key blob,
    column1 blob,
    value blob,
    PRIMARY KEY ((key), column1)
  ) WITH COMPACT STORAGE;
````

Row key time index

- [ ] TODO: what is it used for
- NOTE: it does not use compact storage
  - [ ] TODO: why ...

````sql
CREATE TABLE IF NOT EXISTS row_key_time_index (
    metric text,
    row_time timestamp,
    value text,
    PRIMARY KEY ((metric), row_time)
  )
````

Row keys

- NOTE: no compact storage
- [ ] TODO: metric is text instead of blob
- [ ] TODO: what is frozen<map<text, text>>
- [ ] TODO: compound primary key? http://docs.datastax.com/en/cql/3.3/cql/cql_using/useCompoundPrimaryKey.html
- [ ] TODO: composite partition key https://docs.datastax.com/en/cql/3.1/cql/cql_reference/refCompositePk.html

````sql
CREATE TABLE IF NOT EXISTS row_keys (
    metric text,
    row_time timestamp,
    data_type text,
    tags frozen<map<text, text>>,
    value text,
    PRIMARY KEY ((metric, row_time), data_type, tags)
  )
````

String index

- [ ] TODO: what is it used for
- [ ] TODO: why they are similar, string_index, row_key_index, data_points

````sql
CREATE TABLE IF NOT EXISTS string_index (
    key blob,
    column1 blob,
    value blob,
    PRIMARY KEY ((key), column1)
  ) WITH COMPACT STORAGE
````

### Meta Schema

### Extra Schema
