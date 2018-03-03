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


## Other databases

### InfluxDB