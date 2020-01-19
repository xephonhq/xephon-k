# Xephon-K

<h1 align="center">
	<br>
	<img width="400" src="https://raw.githubusercontent.com/at15/artwork/master/logo/xephonhq/xephon-k.png" alt="xephon-k">
	<br>
	<br>
	<br>
</h1>

[![GoDoc](https://godoc.org/github.com/xephonhq/xephon-k?status.svg)](https://godoc.org/github.com/xephonhq/xephon-k)
[![Build Status](https://travis-ci.org/xephonhq/xephon-k.svg?branch=master)](https://travis-ci.org/xephonhq/xephon-k)
[![Coverage Status](https://coveralls.io/repos/github/xephonhq/xephon-k/badge.svg?branch=master)](https://coveralls.io/github/xephonhq/xephon-k?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/xephonhq/xephon-k)](https://goreportcard.com/report/github.com/xephonhq/xephon-k)
[![codebeat badge](https://codebeat.co/badges/2b3dad97-6550-4b76-a563-a3330d980b23)](https://codebeat.co/projects/github-com-xephonhq-xephon-k-master)

Xephon-K is a time series database with multiple backends. 
It's a playground for comparing modern TSDB design and implementation.
It is not for production use, but it can show you simplified implementation of popular TSDBs. 
A detailed (but not well organized) survey can be found in [doc/survey](doc/survey).

- status: Under major rewrite, after libtsdb-go is stable and benchhub is usable (again), all old code are moved to [_legacy](_legacy)
- [Slide: Xephon-K A Time Series Database with multiple backends](http://www.slideshare.net/ssuser7e134a/xephon-k-a-time-series-database-with-multiple-backends)
- [Survey on existing Cassandra based TSDBs (now include TSDB with other backends)](doc/survey)

## Supported backend

- In Memory
- Local disk, modeled after [InfluxDB](https://github.com/influxdata/influxdb)
- Cassandra, modeled after [KairosDB](https://github.com/kairosdb/kairosdb), but the partitioned schema is not implemented

Following are some backends I plan to implement in the future

- RocksDB
- Redis
- MySQL, modeled after VividCortex
- S3 + Dynamo, modeled after [weaveworks' cortex](https://github.com/weaveworks/cortex/)

## Related projects

- Awesome list [awesome-time-series-database](https://github.com/xephonhq/awesome-time-series-database)
- Benchmark suite [xephon-b](https://github.com/xephonhq/xephon-b) Might still compile

## About

- Xephon comes from animation [RahXephon](https://en.wikipedia.org/wiki/RahXephon), which is chosen for [Xephon-B](https://github.com/xephonhq/xephon-b)
- K comes from KairosDB because this project is originally modeled after KairosDB, which is also the first TSDB I have used in production.

## Authors

- [Pinglei Guo](https://at15.github.io) [@at15](https://github.com/at15)
