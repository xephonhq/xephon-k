# Playground

For testing language semantics and do microbenchmark.
Prototype and library examples are also put here.  

## Language

- https://github.com/golang/go/wiki/SliceTricks

## Prototype

- API: xapi
- Cassandra Bucket: xbucket
  - Not implemented, only create key space
- Cassandra Naive: xnaive 
  - put all data of one series in a single physical row
  - a.k.a use series name as partition key
- Local disk: disk
  - write header, points and part of index
  - can't read them out due to incomplete index
- Local disk C: disk-c
  - only read header