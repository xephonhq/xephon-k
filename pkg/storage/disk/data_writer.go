package disk

/*

On disk file has header, blocks, index, footer

 ---------------------------------------------------
|  Header   |   Blocks   |  Index     |    Footer   |
|  9 bytes  |	N bytes	 |	M bytes	  |   17 bytes  |
 ---------------------------------------------------

- Data that is not serialized (compress, protobuf) all use BigEndian
  - uint32, relative offset and length
  - uint64, absolute offset and magic

- Header:

 -----------------------------------------
| Magic (xephon-k in BigEndian) | Version |
|            8 bytes            |  1 byte |
 -----------------------------------------

- Footer: Index offset  + Index length + Version + Magic
  - [ ] TODO: we don't really need index length, because we can calculate it from offset, but it could be a double check
    - index length = file length - index offset - footer length
  - index length includes index of indexes

 ------------------------------------------------------------------------------------------------
| Index offset | Index of Index offset | Index length |  Version | Magic (xephon-k in BigEndian) |
|   uint64     |        uint32         |    uint32    |          |             uint64            |
|   8 bytes    |        4 byte         |    4 bytes   |  1 byte  |            8 bytes            |
 ------------------------------------------------------------------------------------------------

- Blocks
- [ ] TODO: need offset of values, for Sum, I only need value, don't need time
- [ ] TODO: I think it's better to store the encoding, offset of time and value into index
  - and only let the block store timestamps and values
  - [ ] TODO: or we can store series ID, so we can rebuild index, if file is closed before writing index

 -----------------------------------------------------------------------
| time encoding | value encoding |  encoded timestamps |  encoded values |
|  1 byte       |  1 byte        |      A bytes        |     B bytes     |
 ------------------------------------------------------------------------

- Index
  - Data: Entries,
  	- Entries is `IndexEntries` serialized in protobuf, including series meta and []IndexEntry
  - Index (of Index): series count, series ID, offset, length ...
  	- offset is relative to index begin, not file begin (a.k.a not absolute offset)
  - Footer is the footer of the file, index itself does not have separate footer


 --------------------------------------------------------------------------------------------------------
| Entries1 | Entries2 | ... | series count | series ID1 |  offset  | length  | series ID2 | ... | Footer |
| protobuf | protobuf | ... |   uint32     |   uint64   |  unit32  | uint32  |   uint64   | ... |        |
| A bytes  |  B bytes | ..  |   4 bytes    |  8 bytes   |  4 bytes | 4 bytes |  8 bytes   | ... |        |
 --------------------------------------------------------------------------------------------------------

The writer code is greatly inspired by InfluxDB https://github.com/influxdata/influxdb/blob/master/tsdb/engine/tsm1/writer.go

*/

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"sort"

	"reflect"

	"os"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
)

var _ DataFileWriter = (*LocalDataFileWriter)(nil)
var _ DataFileIndexWriter = (*LocalDataFileIndexWriter)(nil)

var (
	ErrNoData       = fmt.Errorf("no data written, can't write index")
	ErrNotFinalized = fmt.Errorf("index is not written, the file is unreadable")
	ErrFinalized    = fmt.Errorf("index is already written, this file can no longer be updated")
)

const (
	DefaultBufferSize      = 4 * 1024  // 4KB, same as bufio defaultBufSize, InfluxDB use 1MB
	IndexOfIndexUnitLength = 8 + 4 + 4 // id + offset + length
	FooterLength           = 25
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
	// TODO: use must means we will panic if use non exist ID
	MustEntries(id common.SeriesID) *IndexEntries
	WriteAll(io.Writer) (length uint32, indexOffset uint32, errs error)
}

type LocalDataFileWriter struct {
	f         *os.File
	w         *bufio.Writer
	index     DataFileIndexWriter
	n         uint64
	finalized bool
}

type LocalDataFileIndexWriter struct {
	series map[common.SeriesID]*IndexEntries
}

//func NewLocalFileWriter(w io.WriteCloser, bufferSize int) *LocalDataFileWriter {
func NewLocalFileWriter(f *os.File, bufferSize int) (*LocalDataFileWriter, error) {
	if bufferSize <= 0 {
		bufferSize = DefaultBufferSize
	}

	return &LocalDataFileWriter{
		f:         f,
		w:         bufio.NewWriterSize(f, bufferSize),
		index:     NewLocalFileIndexWriter(),
		n:         0,
		finalized: false,
	}, nil
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
	var (
		tenc                     TimeEncoder
		venc                     ValueEncoder
		tBytes, vBytes           []byte
		tBytesCount, vBytesCount int
		err                      error
	)

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

	// write index
	indexOffset := writer.n
	indexLength, indexOfIndexOffset, err := writer.index.WriteAll(writer.w)
	if err != nil {
		return errors.Wrap(err, "can't write index")
	}

	// write footer
	// | index offset (8) | index of index offset (4) | index length (4) | version (1) | magic (8) |
	var buf [FooterLength]byte
	binary.BigEndian.PutUint64(buf[:8], indexOffset)
	binary.BigEndian.PutUint32(buf[8:12], indexOfIndexOffset)
	binary.BigEndian.PutUint32(buf[12:16], indexLength)

	buf[16] = Version
	binary.BigEndian.PutUint64(buf[17:], MagicNumber)
	n, err := writer.w.Write(buf[:])
	if err != nil {
		return errors.Wrap(err, "can't write index position and magic number")
	}
	if n != FooterLength {
		return errors.Errorf("footer should be %d bytes, but %d is written", FooterLength, n)
	}
	writer.finalized = true
	return nil
}

func (writer *LocalDataFileWriter) Flush() error {
	if err := writer.w.Flush(); err != nil {
		return errors.Wrap(err, "can't flush bufio.Writer")
	}

	//if f, ok := writer.f.(*os.File); ok {
	if err := writer.f.Sync(); err != nil {
		return errors.Wrap(err, "can't flush underlying os.File")
	}
	//}
	return nil
}

func (writer *LocalDataFileWriter) Close() error {
	if !writer.Finalized() {
		return ErrNotFinalized
	}
	if err := writer.Flush(); err != nil {
		return errors.Wrap(err, "can't flush before close")
	}
	if err := writer.f.Close(); err != nil {
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
			SeriesMeta: series.GetMetaCopy(),
		}
		idx.series[id] = entries
	}
	entries.Entries = append(entries.Entries, IndexEntry{
		Offset:    offset,
		BlockSize: size,
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

func (idx *LocalDataFileIndexWriter) WriteAll(w io.Writer) (length uint32, indexOffset uint32, errs error) {
	N := 0
	ids := idx.SortedID()
	// index of indexes, written at the last of index
	// | count (4) | series ID (8) | offset (4) | length (4) |
	index := make([]byte, 4+IndexOfIndexUnitLength*len(ids))
	binary.BigEndian.PutUint32(index[:4], uint32(len(ids)))

	for i, id := range ids {
		// TODO: InfluxDB sort the index entry in entries by time before write, but it's likely the blocks of one series is written in time order
		entries := idx.series[id]
		log.Tracef("write: IndexEntries %s", entries)
		b, err := entries.Marshal()
		log.Trace("write: full bytes of IndexEntries")
		log.Trace(b)
		if err != nil {
			errs = errors.Wrap(err, "can't marshal IndexEntries using protobuf")
			return
		}
		n, err := w.Write(b)
		if err != nil {
			length = uint32(N + n)
			errs = errors.Wrap(err, "can't write marshaled IndexEntries to writer")
			return
		}

		start := 4 + i*IndexOfIndexUnitLength
		// id
		binary.BigEndian.PutUint64(index[start:start+8], uint64(id))
		log.Tracef("write: id %d", id)
		// offset
		binary.BigEndian.PutUint32(index[start+8:start+12], uint32(N))
		log.Tracef("write: index offset %d", N)
		// length
		binary.BigEndian.PutUint32(index[start+12:start+16], uint32(n))
		log.Tracef("write: index length %d", n)

		N += n
	}

	log.Trace("write: full bytes for index of indexes")
	log.Trace(index)

	// write index of indexes
	n, err := w.Write(index)
	length = uint32(N + n)
	if err != nil {
		errs = errors.Wrap(err, "cant write index of indexes to writer")
		return
	}
	if n != len(index) {
		errs = errors.Errorf("index of indexes should be %d bytes, but %d written", len(index), n)
	}

	indexOffset = uint32(N)
	errs = nil
	return
}
