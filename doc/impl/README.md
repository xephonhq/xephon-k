# Implementation

A general overview of the implementation. Highly likely to be outdated.

## Write

- client
- HTTP Server `pkg/server/http_server.go`
- JSON decoded `pkg/server/service/write.go` `MakeDecode`
  - use `MetaSeries` to decode meta first (TODO: this might introduce extra copy of the points part in payload)
  - switch `SeriesType` to decode into `IntSeries` and `DoubleSeries`
    - FIXME: return error if meet un matched `SeriesType`, (should allow others to pass?)
- EndPoint `pkg/server/service/write.go` `MakeEndpoint`
    - `WriteServiceServerImpl` call `WriteInt` of its `Store` (memory, disk, cassandra etc.)
- Store `pkg/storage/store.go` defines the interface
  - `pkg/storage/memory/store.go` `WriteIntSeries` write both the data and index
  - `pkg/storage/cassandra/store.go` `WriteIntSeries`  only write data
  - `pkg/storage/disk/store.go` TODO: nothing is implemented
- Memory: WriteIntSeries Data
  - `pkg/storage/memory/data.go` `WriteIntSeries` write to different maps for different types
  - TODO: this introduce problem for lookup, should use a single map, otherwise for each series ID we need two lookup
  - the two maps contain pointer to `IntSeriesStore`, `DoubleSeriesStore` separately.
- Memory: IntSeriesStore
  - `pkg/storage/memory/series_store.go` protect a `common.IntSeries` with a RWMutex
  - TODO: should use several (link list) blocks instead of a huge slice
  - `WriteSeries` sort the input and do a merge sort
    - TODO: the sort can be avoided if we allow user to specify it when input
    - TODO: there are many possible speed ups, like append it directly
    - TODO: compute the common aggregation (allow user to configure it)
- Memory: WriteIntSeries Index
  - `pkg/storage/memory/index.go` `Index Add(id, tagKey, tagValue)` 
    - create new inverted index if `tagKeytagValue` does not exists FIXME: currently direct concat the two string
    - call `InvertedIndex Add(id)`, use quick search to find if the id already exists, otherwise the position to insert it (to keep the order)
    - TODO: could add a map as extra layer to speed up insert and merge it in in batch

## Read

TBD