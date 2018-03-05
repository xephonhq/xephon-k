# In memory data store

- issue: https://github.com/xephonhq/xephon-k/issues/65

## Previously, on Xephon-K

Old in memory data store, just a map with mutex,

````go
// Data is a map using SeriesID as key
type Data struct {
	mu     sync.Mutex
	series map[common.SeriesID]SeriesStore
}
````

## Problems

- reduce lock contention, need benchmark to see what happens to map with a lock with many go routine try to access it
  - #57 https://github.com/xephonhq/xephon-k/issues/57 storage memory concurrent map write
  - sharding could be a way
  - [ ] TODO: see how prometheus and influxdb handle this
- reduce gc overhead
  - [ ] https://github.com/allegro/bigcache https://allegro.tech/2016/03/writing-fast-cache-service-in-go.html
- test with larger memory usage


## Go specific

- https://github.com/allegro/bigcache
- https://github.com/coocood/freecache

## Other databases

### InfluxDB

- see [doc/survey/influxdb/write-path.md](../survey/influxdb/write-path.md) and [doc/survey/influxdb/read-path.md](../survey/influxdb/read-path.md)
- use a ring shard (its robin hood hashing is used for meta index), it has # of cores partitions
- each partition has a `map[string]*entry` where key is measurement + tags + field key, partition has rwmutex
- each entry contains a slice of values, entry has rwmutex
- data in memory is not compressed
- default cache max memory is 1GB 
- default cache snapshot memory size is 25MB 
- default cache snapshot duration is 10 minute