# Heroic

- GitHub: https://github.com/spotify/heroic

## Protocol

- CQL

## Schema

- [legacy](https://github.com/spotify/heroic/tree/master/metric/datastax/src/main/resources/com.spotify.heroic.metric.datastax.schema.legacy)
- [next generation](https://github.com/spotify/heroic/tree/master/metric/datastax/src/main/resources/com.spotify.heroic.metric.datastax.schema.ng)

### Metrics Schema

Keyspace

````
CREATE KEYSPACE IF NOT EXISTS {{keyspace}}
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 3
  };
````

````
CREATE TABLE IF NOT EXISTS {{keyspace}}.metrics (
  metric_key blob,
  data_timestamp_offset int,
  data_value double,
  PRIMARY KEY(metric_key, data_timestamp_offset)
) WITH COMPACT STORAGE;
````

### Meta Schema

### Extra Schema
