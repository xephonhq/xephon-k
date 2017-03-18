# Schema Specification Draft

## Naive

Naive schema pretty looks like [Heroic's schema](../survey/heroic.md), except their row key is actually
a mix of metric name of tags, I don't know if they have bucket (time range) in the
row key

- [implementation](../pkg/bin/xnaive/main.go)

### Requirement

- don't consider tags
- don't consider the size of a phsyical row
- only support int as value

### DDL

Keyspace

````sql
CREATE KEYSPACE IF NOT EXISTS "xephonnaive"
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
  };
````

Metrics

````sql
CREATE TABLE IF NOT EXISTS "xephonnaive".metrics (
  metric_name text,
  metric_timestamp timestamp,
  value int,
  PRIMARY KEY (metric_name, metric_timestamp)
)
````

- `create keyspace if not exists "xephonnaive" with replication = {'class': 'SimpleStrategy', 'replication_factor': 1}; `
- `create table if not exists "xephonnaive".metrics (metric_name text, metric_timestamp timestamp, value int, PRIMARY KEY (metric_name, metric_timestamp));`

### DML

**NOTE: can NOT use double quote for string values!**

Insert

````sql
INSERT INTO "xephonnaive".metrics (metric_name, metric_timestamp, value)
  VALUES ('cpu.load', now(), 30)
````

Select

````sql
SELECT * FROM "xephonnaive".metrics
````

- `insert into "xephonnaive".metrics (metric_name, metric_timestamp, value) values ("cpu.load", now(), 30);`
  - [x] FIXME: `SyntaxException: line 1:91 no viable alternative at input ',' (... metric_timestamp, value) values (["cpu.loa]d",...)`
  - NOTE: can NOT use double quote for string values!
  - https://docs.datastax.com/en/cql/3.3/cql/cql_reference/escape_char_r.html
  - > Column names that contain characters that CQL cannot parse need to be enclosed in double quotation marks in CQL
  - > Dates, IP addresses, and strings need to be enclosed in single quotation marks. To use a single quotation mark itself in a string literal, escape it using a single quotation mark.
- `select * from "xephonnaive".metrics`
  - `Cannot execute this query as it might involve data filtering and thus may have unpredictable performance.If you want to execute this query despite the performance unpredictability, use ALLOW FILTERING`
  - http://stackoverflow.com/questions/38350656/cassandra-asks-for-allow-filtering-even-though-column-is-clustering-key
  - need to specify row key
  - `2017/02/05 11:11:45 gocql: not enough columns to scan into: have 2 want 3` when using `SELECT *`
  - need to have enough variables when use `iter.Scan`

## Bucket

Keyspace and Metrics table remain the same, the main is bucket

- the naive schema simply use user provided name as `metric_name` and as row key (partition key),
which won't scale, we now add bucket to ~~`metric_name`~~ row key
- the bucket value is the start of time of this physical row, i.e. `cpu-load-1487315000`, in other word **each physical row is a bucket**
  - I think we can use two column to be the row key instead of concat them in the tsdb side
- use fixed bucket size, though using dynmaic one would be more efficient, but it requires more work
- [ ] when using bucket, do we still store full timestamp or we just store the delta
  - [ ] this may differ the logic of aggregation and if we can use cassandra's aggregation operations

Keyspace

````sql
CREATE KEYSPACE IF NOT EXISTS "xephonbucket"
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
  };
````

Metrics

````sql
CREATE TABLE IF NOT EXISTS "xephonbucket".metrics (
  metric_name text,
  bucket timestamp,
  metric_timestamp timestamp,
  value int,
  PRIMARY KEY ((metric_name, bucket), metric_timestamp)
)
````

Meta table is used for the following

- keep track of buckets, so if someone only provide one side of the range (only start or only end), you can specify the right row keys to looking for
  - tsdb node read meta table when start up, create the table if not exists
  - a separated go routine to write the bucket name to meta table
    - [ ] there is a problem that newest bucket won't be read because the meta table is not updated immediately
      - maybe notify the go routine when a new bucket is created
- [ ] TODO: I think there is more, not tags, but I went to play life is strange after I thought about it....
- [ ] TODO: maybe I can also save bucket size in the metrics table into the row key

I think it's possible to add bucket now, just add a map into the row key part would work I guess, and still
need meta table to speed up I guess

Meta

````sql
CREATE TABLE IF NOT EXISTS "xephonbucket".meta (
  metric_name text,
  bucket timestamp,
  size int,
  PRIMARY KEY (metric_name, bucket)
)
````

## Naive with Tags

Schema

````sql
CREATE KEYSPACE IF NOT EXISTS "xephonnaivetag"
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
  };
````

Metrics

````sql
CREATE TABLE IF NOT EXISTS "xephonnaivetag".metrics (
  metric_name text,
  metric_timestamp timestamp,
  tags frozen<map<text, text>>,
  value int,
  PRIMARY KEY (metric_name, metric_timestamp, tags)
)
````
