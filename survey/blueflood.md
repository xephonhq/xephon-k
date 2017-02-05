# Blueflood

- GitHub: https://github.com/rackerlabs/blueflood/

## Protocol

## Schema

- https://github.com/rackerlabs/blueflood/blob/master/src/cassandra/cli/load.cdl
- it has separated rollup table

### Metrics Schema

full, 5m, 20m, 60m, 240m, 1440m,

````sql
CREATE TABLE IF NOT EXISTS "DATA".metrics_full (
    key text,
    column1 bigint,
    value blob,
    PRIMARY KEY (key, column1)
) WITH COMPACT STORAGE AND speculative_retry = 'NONE';

CREATE TABLE IF NOT EXISTS "DATA".metrics_5m (
    key text,
    column1 bigint,
    value blob,
    PRIMARY KEY (key, column1)
) WITH COMPACT STORAGE AND speculative_retry = 'NONE';
````

preaggreated full, 5m, 20m, 60m, 240m, 1440m,

````sql
CREATE TABLE IF NOT EXISTS "DATA".metrics_preaggregated_full (
    key text,
    column1 bigint,
    value blob,
    PRIMARY KEY (key, column1)
) WITH COMPACT STORAGE AND speculative_retry = 'NONE';

CREATE TABLE IF NOT EXISTS "DATA".metrics_preaggregated_5m (
    key text,
    column1 bigint,
    value blob,
    PRIMARY KEY (key, column1)
) WITH COMPACT STORAGE AND speculative_retry = 'NONE';
````

### Meta Schema

````sql
CREATE TABLE IF NOT EXISTS "DATA".metrics_metadata (
    key text,
    column1 text,
    value blob,
    PRIMARY KEY (key, column1)
) WITH COMPACT STORAGE AND speculative_retry = 'NONE';
````

### Extra Schema
