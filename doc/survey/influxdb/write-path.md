# InfluxDB Write Path

Copied from libtsdb-go

- tsdb/engine/tsm1/engine.go `func (e *Engine) WritePoints(points []models.Point)`
  - `err := e.Cache.WriteMulti(values)` it writes to cache before write to WAL
  - `_, err = e.WAL.WriteMulti(values)`
- tsdb/engine/tsm1/cache.go use `storer` interface
  - `guru -scope . implements cache.go:#4432` 4432 is byte offset ... not line number, which works for editors, but not human ...
  - implemented by `ring`, `TestStore`, `emptyStore`
- first snapshot, then compactor, them tsmWriter, this is executed async, not sync when request come in
- tsdb/engine/tsm1/engine.go
  - `func (e *Engine) enableSnapshotCompactions()`
  - `func (e *Engine) compactCache()` 'continually checks if the WAL cache should be written to disk'
  - `func (e *Engine) WriteSnapshot() error`
  - `func (e *Engine) writeSnapshotAndCommit(closedFiles []string, snapshot *Cache) (err error)`
- tsdb/engine/tsm1/compact.go
  - `func (c *Compactor) WriteSnapshot(cache *Cache) ([]string, error) {`
  - `func (c *Compactor) writeNewFiles(generation, sequence int, iter KeyIterator, throttle bool) ([]string, error) {`
  - `func (c *Compactor) write(path string, iter KeyIterator, throttle bool) (err error) {`
- tsdb/engine/tsm1/writer.go
  - `func (t *tsmWriter) Write(key []byte, values Values)` saw this function several times when writing Xephon-K(S)
- [] TODO: it calls `Encode` etc. which is where all compression algorithm kicked in
   
````go
// NOTE: some branches and error handling are removed
func (e *Engine) WritePoints(points []models.Point) error {
    values := make(map[string][]Value, len(points))
    var keyBuf []byte
    for _, p := range points {
            keyBuf = append(keyBuf[:0], p.Key()...)
            keyBuf = append(keyBuf, keyFieldSeparator...)
            baseLen = len(keyBuf)
    		iter := p.FieldIterator()
    		t := p.Time().UnixNano()
    		for iter.Next() {
    			keyBuf = append(keyBuf[:baseLen], iter.FieldKey()...)
    			var v Value
    			switch iter.Type() {
    			case models.Float:
    				fv, err := iter.FloatValue()
    				v = NewFloatValue(t, fv)
    			case models.Integer:
    				iv, err := iter.IntegerValue()
    				v = NewIntegerValue(t, iv)
    			default:
    				return fmt.Errorf("unknown field type for %s: %s", string(iter.FieldKey()), p.String())
    			}
    			values[string(keyBuf)] = append(values[string(keyBuf)], v)
    		}
    	}
    }
    // first try to write to the cache
    err := e.Cache.WriteMulti(values)
	_, err = e.WAL.WriteMulti(values)
	return err
}
````

The working cache implementation is ring, in `tsdb/engine/tsm1/ring.go`

- it use a ring for shard, ring itself does not have lock
- each partition has a rw lock and a map of entries, the key of map is measurement + tags + field key 
  - see `keyBuf = append(keyBuf[:baseLen], iter.FieldKey()...)` in `*Engine) WritePoints`, a multi field point is fan out to multiple series at last
- each entry has a rw lock and a slice of value

````go
type ring struct {
	keysHint int64
	partitions []*partition
}

// partition provides safe access to a map of series keys to entries.
type partition struct {
	mu    sync.RWMutex
	store map[string]*entry
}

// entry is a set of values and some metadata.
type entry struct {
	mu     sync.RWMutex
	values Values // All stored values.
	// The type of values stored. Read only so doesn't need to be protected by mu.
	vtype byte
}

type Values []Value

// Value represents a TSM-encoded value.
type Value interface {
	// UnixNano returns the timestamp of the value in nanoseconds since unix epoch.
	UnixNano() int64

	// Value returns the underlying value.
	Value() interface{}

	// Size returns the number of bytes necessary to represent the value and its timestamp.
	Size() int

	// String returns the string representation of the value and its timestamp.
	String() string

	// internalOnly is unexported to ensure implementations of Value
	// can only originate in this package.
	internalOnly()
}

// write writes values to the entry in the ring's partition associated with key.
// If no entry exists for the key then one will be created.
// write is safe for use by multiple goroutines.
func (r *ring) write(key []byte, values Values) (bool, error) {
	return r.getPartition(key).write(key, values)
}

// getPartition retrieves the hash ring partition associated with the provided
// key.
func (r *ring) getPartition(key []byte) *partition {
	return r.partitions[int(xxhash.Sum64(key)%partitions)]
}

// write writes the values to the entry in the partition, creating the entry
// if it does not exist.
// write is safe for use by multiple goroutines.
func (p *partition) write(key []byte, values Values) (bool, error) {
	p.mu.RLock()
	e := p.store[string(key)]
	p.mu.RUnlock()
	if e != nil {
		// Hot path.
		return false, e.add(values)
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	// Check again.
	if e = p.store[string(key)]; e != nil {
		return false, e.add(values)
	}

	// Create a new entry using a preallocated size if we have a hint available.
	e, err := newEntryValues(values)
	if err != nil {
		return false, err
	}

	p.store[string(key)] = e
	return true, nil
}
````