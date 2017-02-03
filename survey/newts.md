# Newts

- GitHub: https://github.com/OpenNMS/newts
- They use Dropwizard for REST API https://github.com/OpenNMS/newts/tree/master/rest

## Protocol

- CQL

## Schema

- [Samples Schema CQL](https://github.com/OpenNMS/newts/blob/master/cassandra/storage/src/main/resources/samples_schema.cql)
  - [x] ~~TODO: is this just the sample, or they are using it?~~
  - It's **samples** not sample ...

Keyspace

````sql
CREATE KEYSPACE $KEYSPACE$
    WITH replication = {'class': 'SimpleStrategy', 'replication_factor': $REPLICATION_FACTOR$};
````

Metrics

- [ ] TODO: the attributes field is the `tags`?
- [ ] TODO: is this a tall and narrow table?
- [ ] what are the `context`, `partition`, `resource`

````sql
CREATE TABLE $KEYSPACE$.samples (
    context text,
    partition int,
    resource text,
    collected_at timestamp,
    metric_name text,
    value blob,
    attributes map<text, text>,
    PRIMARY KEY((context, partition, resource), collected_at, metric_name)
);
````
