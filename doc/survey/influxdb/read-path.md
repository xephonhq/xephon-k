# InfluxDB Read Path

- old query starts from `services/httpd/handler.go`

- new query (using ifql) starts from `services/storage` using yarpc, see proto file for definition

## Old

- old query starts from `services/httpd/handler.go`
- query/query_executor.go 
  - `func (e *QueryExecutor) ExecuteQuery(query *influxql.Query, opt ExecutionOptions, closing chan struct{}) <-chan *Result`
  - `err = e.StatementExecutor.ExecuteStatement(stmt, ctx)`
- coordinator/statement_executor.go
  - `func (e *StatementExecutor) ExecuteStatement(stmt influxql.Statement, ctx query.ExecutionContext) error`
  - `func (e *StatementExecutor) executeSelectStatement(ctx context.Context, stmt *influxql.SelectStatement, ectx *query.ExecutionContext) error`
    - `itrs, stmt, err := e.createIterators(stmt, ctx)`
- [ ] TODO: lost in the middle
- tsdb/engine.go
  - Engine interface `CreateIterator(ctx context.Context, measurement string, opt query.IteratorOptions) (query.Iterator, error)`
    - it also has trace /w\
- tsdb/engine/tsm1/engine.go
  - `func (e *Engine) CreateIterator(ctx context.Context, measurement string, opt query.IteratorOptions) (query.Iterator, error)`
  - `func (e *Engine) createVarRefIterator(ctx context.Context, measurement string, opt query.IteratorOptions) ([]query.Iterator, error)`
   - [ ] TODO: lost in the middle
   - cursor is used in iterator I think ...
  - `Cache` is ued for building batch cursor, i.e. `func (e *Engine) buildBooleanBatchCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) tsdb.BooleanBatchCursor {`
    - `cacheValues := e.Cache.Values(key)`
- tsdb/engine/tsm1/cache.go
  - `func (c *Cache) Values(key []byte) Values {` 'Values returns a copy of all values, deduped and sorted, for the given key.' 
- tsdb/engine/tsm1/file_store.go
  - `KeyCursor` is used for reading file
  - `ReadBooleanBlock` (other types are also generated from template) is called to read data into the cursor
  - where `ReadBooleanBlockAt` is called 
- tsdb/engine/tsm1/reader.go
  - `func (m *mmapAccessor) readBooleanBlock(entry *IndexEntry, values *[]BooleanValue) ([]BooleanValue, error) {`
- tsdb/engine/tsm1/encoding.go
  - `func DecodeBooleanBlock(block []byte, a *[]BooleanValue) ([]BooleanValue, error) {`
- tsdb/engine/tsm1/timestamp.go for timestamp encoder and decoder
- tsdb/engine/tsm1/float.go for float decoder etc.

````go
// buildBooleanBatchCursor creates a batch cursor for a boolean field.
func (e *Engine) buildBooleanBatchCursor(ctx context.Context, measurement, seriesKey, field string, opt query.IteratorOptions) tsdb.BooleanBatchCursor {
	key := SeriesFieldKeyBytes(seriesKey, field)
	cacheValues := e.Cache.Values(key)
	keyCursor := e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	return newBooleanBatchCursor(seriesKey, opt.SeekTime(), opt.Ascending, cacheValues, keyCursor)
}


// KeyCursor returns a KeyCursor for the given key starting at time t.
func (e *Engine) KeyCursor(ctx context.Context, key []byte, t int64, ascending bool) *KeyCursor {
	return e.FileStore.KeyCursor(ctx, key, t, ascending)
}

// FileStore is an abstraction around multiple TSM files.
type FileStore struct {
	  mu    sync.RWMutex
	  files []TSMFile
}

// KeyCursor returns a KeyCursor for key and t across the files in the FileStore.
func (f *FileStore) KeyCursor(ctx context.Context, key []byte, t int64, ascending bool) *KeyCursor {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return newKeyCursor(ctx, f, key, t, ascending)
}

// TSMFile represents an on-disk TSM file.
type TSMFile interface {
	  // ReadAt returns all the values in the block identified by entry.
  	ReadAt(entry *IndexEntry, values []Value) ([]Value, error)
  	ReadFloatBlockAt(entry *IndexEntry, values *[]FloatValue) ([]FloatValue, error)
  	ReadIntegerBlockAt(entry *IndexEntry, values *[]IntegerValue) ([]IntegerValue, error)
  	ReadUnsignedBlockAt(entry *IndexEntry, values *[]UnsignedValue) ([]UnsignedValue, error)
  	ReadStringBlockAt(entry *IndexEntry, values *[]StringValue) ([]StringValue, error)
  	ReadBooleanBlockAt(entry *IndexEntry, values *[]BooleanValue) 
}

func (c *KeyCursor) ReadBooleanBlock(buf *[]BooleanValue) ([]BooleanValue, error) {
LOOP:
	// No matching blocks to decode
	if len(c.current) == 0 {
		return nil, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	*buf = (*buf)[:0]
	values, err := first.r.ReadBooleanBlockAt(&first.entry, buf)
	if err != nil {
		return nil, err
	}
	// ... 
}

func (m *mmapAccessor) readBooleanBlock(entry *IndexEntry, values *[]BooleanValue) ([]BooleanValue, error) {
	m.incAccess()

	m.mu.RLock()
	if int64(len(m.b)) < entry.Offset+int64(entry.Size) {
		m.mu.RUnlock()
		return nil, ErrTSMClosed
	}

	a, err := DecodeBooleanBlock(m.b[entry.Offset+4:entry.Offset+int64(entry.Size)], values)
	m.mu.RUnlock()

	if err != nil {
		return nil, err
	}

	return a, nil
}

// DecodeBooleanBlock decodes the boolean block from the byte slice
// and appends the boolean values to a.
func DecodeBooleanBlock(block []byte, a *[]BooleanValue) ([]BooleanValue, error) {
	// Block type is the next block, make sure we actually have a float block
	blockType := block[0]
	if blockType != BlockBoolean {
		return nil, fmt.Errorf("invalid block type: exp %d, got %d", BlockBoolean, blockType)
	}
	block = block[1:]

	tb, vb, err := unpackBlock(block)
	if err != nil {
		return nil, err
	}

	sz := CountTimestamps(tb)

	if cap(*a) < sz {
		*a = make([]BooleanValue, sz)
	} else {
		*a = (*a)[:sz]
	}
	tdec := timeDecoderPool.Get(0).(*TimeDecoder)
  vdec := booleanDecoderPool.Get(0).(*BooleanDecoder)
 // ...
}


// TimeDecoder decodes a byte slice into timestamps.
type TimeDecoder struct {
	v    int64
	i, n int
	ts   []uint64
	dec  simple8b.Decoder
	err  error

	// The delta value for a run-length encoded byte slice
	rleDelta int64

	encoding byte
}

// FloatEncoder encodes multiple float64s into a byte slice.
type FloatEncoder struct {
	val float64
	err error

	leading  uint64
	trailing uint64

	buf bytes.Buffer
	bw  *bitstream.BitWriter

	first    bool
	finished bool
}

````