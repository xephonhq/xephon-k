# xephon-k

[![GoDoc](https://godoc.org/github.com/xephonhq/xephon-k?status.svg)](https://godoc.org/github.com/xephonhq/xephon-k)
[![Go Report Card](https://goreportcard.com/badge/github.com/xephonhq/xephon-k)](https://goreportcard.com/report/github.com/xephonhq/xephon-k)

A time series database using Cassandra as backend, modeled after KairosDB

## Roadmap

- [Survey](survey)
  - [x] [existing TSDBs using C*](https://github.com/xephonhq/awesome-time-series-database#cassandra)
  - [x] categorize schemas
- [Specification](doc/spec-draft.md)
  - [x] naive schema
  - [x] [naive schema's naive implementation](pkg/bin/xnaive/main.go)
- storing metrics
  - [x] memory without index tag
  - [x] cassandra without index tag
- query metrics as it is
  - [x] without using tag
- query aggregation
- index text without using external search engine
  - i.e. https://github.com/balzaczyy/golucene
