# Khronus


- GitHub: https://github.com/Searchlight/khronus
- It uses https://github.com/HdrHistogram/HdrHistogram for Histogram

## Protocol

- CQL

## Schema

- `create keyspace if not exists $keyspacePlusSuffix with replication = {'class':'SimpleStrategy', 'replication_factor': $getRF};`
- https://github.com/Searchlight/khronus/blob/master/khronus-core/src/main/scala/com/searchlight/khronus/store/Cassandra.scala#L76

### Metrics Schema

- `create table if not exists ${tableName(window)} (metric text, timestamp bigint, buckets ${getBucketsCollectionType(window)}<blob>, primary key (metric, timestamp)) with gc_grace_seconds = 0 and compaction = {'class': 'LeveledCompactionStrategy' }`
- https://github.com/Searchlight/khronus/blob/master/khronus-core/src/main/scala/com/searchlight/khronus/store/BucketStore.scala#L69

### Meta Schema

- `create table if not exists meta (key text, metric text, timestamp bigint, active boolean, primary key (key, metric))`
- https://github.com/Searchlight/khronus/blob/master/khronus-core/src/main/scala/com/searchlight/khronus/store/MetaStore.scala#L65

### Extra Schema

- summary store (seems to be rollup or preaggreate)
  - https://github.com/Searchlight/khronus/blob/master/khronus-core/src/main/scala/com/searchlight/khronus/store/SummaryStore.scala
