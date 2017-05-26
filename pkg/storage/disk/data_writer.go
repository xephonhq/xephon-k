package disk

/*

On disk file has header, blocks, index, footer

 ---------------------------------------------------
|  Header   |   Blocks   |  Index     |    Footer   |
|  9 bytes  |	N bytes	 |	M bytes	  |   17 bytes  |
 ---------------------------------------------------

- Header: Magic (xephon-k in bigendian) + Version
- Footer: Index offset (bigendian uint64) + Version + Magic
- Blocks

| time encoding | value encoding |  encoded timestamps |  encoded values |
|  1 byte       |  1 byte        |      A bytes        |     B bytes     |

The writer code is greatly inspired by InfluxDB https://github.com/influxdata/influxdb/blob/master/tsdb/engine/tsm1/writer.go

*/

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"sort"

	"reflect"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	"os"
)

var _ DataFileWriter = (*LocalDataFileWriter)(nil)
var _ DataFileIndexWriter = (*LocalDataFileIndexWriter)(nil)

var (
	ErrNoData       = fmt.Errorf("no data written, can't write index")
	ErrNotFinalized = fmt.Errorf("index is not written, the file is unreadable")
	ErrFinalized    = fmt.Errorf("index is already written, this file can no longer be updated")
)

const (
	DefaultBufferSize = 4 * 1024 // 4KB, same as bufio defaultBufSize, InfluxDB use 1MB
)

// DataFileWriter writes data to disk, index is at the end of the file for locating data blocks.
// It is NOT thread safe
type DataFileWriter interface {
	// WriteHeader writes the magic number and version, it will be called by WriteSeries automatically for once
	WriteHeader() error
	WriteSeries(series common.Series) error
	// Finalized returns if the index is written and the file can be closed
	Finalized() bool
	// WriteIndex writes index data and trailing magic number into the end of the file (buffer). You can only call it once
	WriteIndex() error
	// Flush flushes data in the buffer to disk
	Flush() error
	// Close closes the underlying file, the file must be finalized, otherwise it can't be read
	Close() error
}

type DataFileIndexWriter interface {
	Add(series common.Series, offset uint64, size uint64) error
	SortedID() []common.SeriesID
	Len() int
	// TODO: we need to record information like int/double, precision, min, max time etc.
	// TODO: we don't need to duplicate series info like name tags for each entry, store in IndexEntries
	// TODO: use must means we will panic if use non exist ID
	MustEntries(id common.SeriesID) *IndexEntries
}

type LocalDataFileWriter struct {
	originalWriter io.WriteCloser
	w              *bufio.Writer
	index          DataFileIndexWriter
	n              uint64
	finalized      bool
}

type LocalDataFileIndexWriter struct {
	series map[common.SeriesID]*IndexEntries
}

func NewLocalFileWriter(w io.WriteCloser, bufferSize int) *LocalDataFileWriter {
	if bufferSize <= 0 {
		bufferSize = DefaultBufferSize
	}

	return &LocalDataFileWriter{
		originalWriter: w,
		w:              bufio.NewWriterSize(w, bufferSize),
		index:          NewLocalFileIndexWriter(),
		n:              0,
		finalized:      false,
	}
}

func NewLocalFileIndexWriter() *LocalDataFileIndexWriter {
	return &LocalDataFileIndexWriter{
		series: map[common.SeriesID]*IndexEntries{},
	}
}

func (writer *LocalDataFileWriter) WriteHeader() error {
	var buf [9]byte
	binary.BigEndian.PutUint64(buf[:8], MagicNumber)
	buf[8] = Version
	n, err := writer.w.Write(buf[:])
	if err != nil {
		return errors.Wrap(err, "can't write header (magic number + version)")
	}
	if n != 9 {
		return errors.Errorf("header should be 9 bytes, but %d is written", n)
	}
	writer.n += 9
	return nil
}

func (writer *LocalDataFileWriter) WriteSeries(series common.Series) error {
	if writer.Finalized() {
		return ErrFinalized
	}
	// write header if this is the first series
	if writer.n == 0 {
		if err := writer.WriteHeader(); err != nil {
			return err
		}
	}

	n := 0
	var tenc TimeEncoder
	var venc ValueEncoder
	var tBytes, vBytes []byte
	var tBytesCount, vBytesCount int
	var err error

	// encode time and value separately
	// TODO: only use RawBigEndianTime/IntEncoder for now
	tenc = &RawBigEndianTimeEncoder{}
	switch series.GetSeriesType() {
	case common.TypeIntSeries:
		intSeries, ok := series.(*common.IntSeries)
		if !ok {
			return errors.Errorf("%s %v is marked as int but actually %s",
				series.GetName(), series.GetTags(), reflect.TypeOf(series))
		}
		intEnc := &RawBigEndianIntEncoder{}
		for i := 0; i < len(intSeries.Points); i++ {
			tenc.Write(intSeries.Points[i].T)
			intEnc.Write(intSeries.Points[i].V)
		}
		venc = intEnc
	default:
		return errors.Errorf("unsupported series type %d", series.GetSeriesType())
	}
	// write encoding information
	writer.w.Write([]byte{tenc.Encoding(), venc.Encoding()})
	n += 2
	// write encoded time and values
	if tBytes, err = tenc.Bytes(); err != nil {
		return errors.Wrap(err, "can't get encoded time as bytes")
	}
	if vBytes, err = venc.Bytes(); err != nil {
		return errors.Wrap(err, "can't get encoded value as bytes")
	}
	if tBytesCount, err = writer.w.Write(tBytes); err != nil {
		return errors.Wrap(err, "cant write encoded time to buffer")
	}
	n += tBytesCount
	if vBytesCount, err = writer.w.Write(vBytes); err != nil {
		return errors.Wrap(err, "can't write encoded value to buffer")
	}
	n += vBytesCount

	// record block position in index
	// TODO: should store some aggregated information as well
	writer.index.Add(series, writer.n, uint64(n))

	writer.n += uint64(n)
	return nil
}

func (writer *LocalDataFileWriter) Finalized() bool {
	return writer.finalized
}

func (writer *LocalDataFileWriter) WriteIndex() error {
	if writer.Finalized() {
		return ErrFinalized
	}

	if writer.index.Len() == 0 {
		return ErrNoData
	}

	// TODO: write index to underlying writer
	indexPos := writer.n

	// write footer
	// | index position (8) | version (1) | magic (8) |
	var buf [17]byte
	binary.BigEndian.PutUint64(buf[:8], indexPos)
	buf[8] = Version
	binary.BigEndian.PutUint64(buf[9:], MagicNumber)
	n, err := writer.w.Write(buf[:])
	if err != nil {
		return errors.Wrap(err, "can't write index position and magic number")
	}
	if n != 17 {
		return errors.Errorf("footer should be 17 bytes, but %d is written", n)
	}
	writer.finalized = true
	return nil
}

func (writer *LocalDataFileWriter) Flush() error {
	if err := writer.w.Flush(); err != nil {
		return errors.Wrap(err, "can't flush bufio.Writer")
	}

	if f, ok := writer.originalWriter.(*os.File); ok {
		if err := f.Sync(); err != nil {
			return errors.Wrap(err, "can't flush underlying os.File")
		}
	}
	return nil
}

func (writer *LocalDataFileWriter) Close() error {
	if !writer.Finalized() {
		return ErrNotFinalized
	}
	if err := writer.Flush(); err != nil {
		return errors.Wrap(err, "can't flush before close")
	}
	if err := writer.originalWriter.Close(); err != nil {
		return errors.Wrap(err, "flushed but can't close")
	}
	return nil
}

func (idx *LocalDataFileIndexWriter) Add(series common.Series, offset uint64, size uint64) error {
	id := series.GetSeriesID()
	entries, ok := idx.series[id]
	// create new IndexEntries if this series has not been added
	if !ok {
		entries = &IndexEntries{
			meta: series.GetMetaCopy(),
		}
		idx.series[id] = entries
	}
	entries.entries = append(entries.entries, IndexEntry{
		Offset: offset,
		Size:   size,
	})
	return nil
}

func (idx *LocalDataFileIndexWriter) SortedID() []common.SeriesID {
	keys := make([]common.SeriesID, 0, len(idx.series))
	for k := range idx.series {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func (idx *LocalDataFileIndexWriter) Len() int {
	return len(idx.series)
}

func (idx *LocalDataFileIndexWriter) MustEntries(id common.SeriesID) *IndexEntries {
	entries, ok := idx.series[id]
	if !ok {
		log.Panicf("can't find entries for %d", id)
	}
	return entries
}
