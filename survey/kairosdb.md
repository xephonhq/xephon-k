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

````
CREATE KEYSPACE IF NOT EXISTS {{keyspace}}
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
  };
````

Metrics

- NOTE: it seems KairosDB has omitted keyspace when creating tables
- NOTE: column1 is not a typo, it might be the legacy problem of using Thrift protocol

````
CREATE TABLE IF NOT EXISTS data_points (
    key blob,
    column1 blob,
    value blob,
    PRIMARY KEY ((key), column1)
  ) WITH COMPACT STORAGE;
````

public static final String DATA_POINTS_TABLE = "" +
    "CREATE TABLE IF NOT EXISTS data_points (\n" +
    "  key blob,\n" +
    "  column1 blob,\n" +
    "  value blob,\n" +
    "  PRIMARY KEY ((key), column1)\n" +
    ") WITH COMPACT STORAGE";

public static final String ROW_KEY_INDEX_TABLE = "" +
    "CREATE TABLE IF NOT EXISTS row_key_index (\n" +
    "  key blob,\n" +
    "  column1 blob,\n" +
    "  value blob,\n" +
    "  PRIMARY KEY ((key), column1)\n" +
    ") WITH COMPACT STORAGE";

public static final String ROW_KEY_TIME_INDEX = "" +
    "CREATE TABLE IF NOT EXISTS row_key_time_index (\n" +
    "  metric text,\n" +
    "  row_time timestamp,\n" +
    "  value text,\n" +
    "  PRIMARY KEY ((metric), row_time)\n" +
    ")";

public static final String ROW_KEYS = "" +
    "CREATE TABLE IF NOT EXISTS row_keys (\n" +
    "  metric text,\n" +
    "  row_time timestamp,\n" +
    "  data_type text,\n" +
    "  tags frozen<map<text, text>>,\n" +
    "  value text,\n" +
    "  PRIMARY KEY ((metric, row_time), data_type, tags)\n" +
    ")";

public static final String STRING_INDEX_TABLE = "" +
    "CREATE TABLE IF NOT EXISTS string_index (\n" +
    "  key blob,\n" +
    "  column1 blob,\n" +
    "  value blob,\n" +
    "  PRIMARY KEY ((key), column1)\n" +
    ") WITH COMPACT STORAGE";
### Meta Schema

### Extra Schema
