# Storage: Memory

Used as a cache to allow batch write to cassandra (and maybe get rid of cassandra in the future)

## Ref 

Ported from [tsdb-proxy-java](https://github.com/xephonhq/tsdb-proxy-java/blob/master/doc/spec-draft/in_memory_store.md)

InfluxDb

- https://github.com/influxdata/influxdb/blob/master/tsdb/engine/tsm1/cache.go
- https://github.com/influxdata/influxdb/blob/master/tsdb/engine/tsm1/ring.go

````go
// NewCache returns an instance of a cache which will use a maximum of maxSize bytes of memory.
// Only used for engine caches, never for snapshots.
func NewCache(maxSize uint64, path string) *Cache {
	store, _ := newring(ringShards)
	c := &Cache{
		maxSize:      maxSize,
		store:        store, // Max size for now..
		stats:        &CacheStatistics{},
		lastSnapshot: time.Now(),
	}
	c.UpdateAge()
	c.UpdateCompactTime(0)
	c.updateCachedBytes(0)
	c.updateMemSize(0)
	c.updateSnapshots()
	return c
}
````

Prometheus

- https://github.com/prometheus/prometheus/blob/master/storage/local/storage.go
- https://github.com/prometheus/prometheus/blob/master/storage/local/series.go

storage.go

- NOTE: they drop data points when they detect duplicate and out of order when append 
 
````go
// Package local contains the local time series storage used by Prometheus.
package local

func (s *MemorySeriesStorage) Append(sample *model.Sample) error {
    if sample.Timestamp == series.lastTime {
		// Don't report "no-op appends", i.e. where timestamp and sample
		// value are the same as for the last append, as they are a
		// common occurrence when using client-side timestamps
		// (e.g. Pushgateway or federation).
		if sample.Timestamp == series.lastTime &&
			series.lastSampleValueSet &&
			sample.Value.Equal(series.lastSampleValue) {
			return nil
		}
		s.discardedSamplesCount.WithLabelValues(duplicateSample).Inc()
		return ErrDuplicateSampleForTimestamp // Caused by the caller.
	}
	if sample.Timestamp < series.lastTime {
		s.discardedSamplesCount.WithLabelValues(outOfOrderTimestamp).Inc()
		return ErrOutOfOrderSample // Caused by the caller.
	}
}
````

````go
package local

// add adds a sample pair to the series. It returns the number of newly
// completed chunks (which are now eligible for persistence).
//
// The caller must have locked the fingerprint of the series.
func (s *memorySeries) add(v model.SamplePair) (int, error) {
	if len(s.chunkDescs) == 0 || s.headChunkClosed {
		newHead := chunk.NewDesc(chunk.New(), v.Timestamp)
		s.chunkDescs = append(s.chunkDescs, newHead)
		s.headChunkClosed = false
	} else if s.headChunkUsedByIterator && s.head().RefCount() > 1 {
		// We only need to clone the head chunk if the current head
		// chunk was used in an iterator at all and if the refCount is
		// still greater than the 1 we always have because the head
		// chunk is not yet persisted. The latter is just an
		// approximation. We will still clone unnecessarily if an older
		// iterator using a previous version of the head chunk is still
		// around and keep the head chunk pinned. We needed to track
		// pins by version of the head chunk, which is probably not
		// worth the effort.
		chunk.Ops.WithLabelValues(chunk.Clone).Inc()
		// No locking needed here because a non-persisted head chunk can
		// not get evicted concurrently.
		s.head().C = s.head().C.Clone()
		s.headChunkUsedByIterator = false
	}

	chunks, err := s.head().Add(v)
	if err != nil {
		return 0, err
	}
	s.head().C = chunks[0]

	for _, c := range chunks[1:] {
		s.chunkDescs = append(s.chunkDescs, chunk.NewDesc(c, c.FirstTime()))
	}

	// Populate lastTime of now-closed chunks.
	for _, cd := range s.chunkDescs[len(s.chunkDescs)-len(chunks) : len(s.chunkDescs)-1] {
		if err := cd.MaybePopulateLastTime(); err != nil {
			return 0, err
		}
	}

	s.lastTime = v.Timestamp
	s.lastSampleValue = v.Value
	s.lastSampleValueSet = true
	return len(chunks) - 1, nil
}
````

Atlas

- https://github.com/Netflix/atlas/wiki/Overview#memory-storage didn't mention how
- [MemoryBlockStore](https://github.com/Netflix/atlas/blob/master/atlas-core/src/main/scala/com/netflix/atlas/core/db/BlockStore.scala)