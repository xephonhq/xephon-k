# Specification Draft

## Naive

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

````sql
INSERT INTO "xephonnaive".metrics (metric_name, metric_timestamp, value)
  VALUES ("cpu.load", now(), 30)
````

- `insert into "xephonnaive".metrics (metric_name, metric_timestamp, value) values ("cpu.load", now(), 30);`
  - [ ] FIXME: `SyntaxException: line 1:91 no viable alternative at input ',' (... metric_timestamp, value) values (["cpu.loa]d",...)`
