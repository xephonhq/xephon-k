# xephon-k

A time series database using Cassandra as backend, modeled after KairosDB

## Roadmap

- [Survey](survey)
  - [x] [existing TSDBs using C*](https://github.com/xephonhq/awesome-time-series-database#cassandra)
  - [ ] categorize schemas
- [Specification](doc/spec-draft.md)
  - [x] naive schema
  - [x] [naive schema's naive implementation](pkg/bin/xnaive/main.go)
- storing metrics
- query metrics as it is
- query aggregation
- index text without using external search engine
  - i.e. https://github.com/balzaczyy/golucene
