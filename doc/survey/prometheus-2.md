# Prometheus

- https://fabxc.org/blog/2017-04-10-writing-a-tsdb/
  - https://github.com/prometheus/tsdb/

## Take away

- https://github.com/prometheus/prombench a E2E benchmarking tool
- http://codecapsule.com/2014/02/12/coding-for-ssds-part-1-introduction-and-table-of-contents/
- Not all time series are long lived, rolling up micro services have burst in new series when update
- A horizontal one https://github.com/weaveworks/cortex

## Writing a Time Series Database from Scratch

V2

- one file per time series that contains all of its samples in sequential order
- batch 1KB chunk in memory and append to individual file
- compression using variant of Gorilla's

Problems of one file per series

- run out of inode
  - only way is to format file system
- thousands of chunks per second
- infeasible to keep all files open for read and write. ~99% data becomes cold
- delete (many) old files further write amplification your SSD
- memory state periodically checkpointed to disk, take several minutes to recover

Series Churn

- [**InfluxDB also addressed this**](influxdb.md)
- some series never get update, i.e. old instances removed by Kubernetes
- leveldb can't handle complex query efficiently i.e. `instance="12bd" AND region="us"`
  - i.e. index tag better

Starting Over

> This is where I take the shortcut and drive straight to the solution — skip the headache, failed ideas, endless sketching, tears, and despair.

- Partition time into non-overlapping blocks
- Each file contains multiple series
- WAL for recent block
- mmap
- compaction of blocks

Retention

- delete directory that contains old blocks

Index

> So assuming our index lookup was of complexity O(n^2), we managed to reduce the n a fair amount and now have an improved complexity of O(n^2) — uhm, wait... damn it.

- [ ] index on tag key or tag value?
- inverted index
- combining labels
  - To find all series satisfying both label selectors, we take the inverted index list for each and intersect them
  - sort the data in inverted index and do a k-way merge `O(nk)`
  - https://github.com/prometheus/tsdb/issues/50 `O(n logk)`
  - https://en.wikipedia.org/wiki/Search_engine_indexing#Inverted_indices

- https://github.com/prometheus/tsdb/blob/master/postings.go in memory operations
- https://github.com/prometheus/tsdb/blob/master/index.go on disk structure

````go
type memPostings struct {
	m map[term][]uint32
}

type term struct {
	name, value string
}
````
