package disk

import (
	"sort"

	"io"

	"bufio"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
)

var _ FileWriter = (*LocalFileWriter)(nil)

// var _ BlockWriter = (*LocalFileBlockWriter)(nil)
var _ IndexWriter = (*LocalFileIndexWriter)(nil)

type FileWriter interface {
	WriteSeries(series common.Series) error
	// Finalize write index data and trailing magic number into the end of the file (buffer). You can only call it once
	// TODO: did/how influxdb restrict it only get called once, and can't write series once index is written
	Finalize() error
	// Flush flushes data in the buffer to disk
	Flush() error
	// Close closes the underlying file, the file must be finalized, otherwise it can't be read
	Close() error
}

// type BlockWriter interface {
// 	WriteSeries(w io.Writer, series common.Series) error
// }

type IndexWriter interface {
	// TODO: why influx db store offset as int64 instead of uint64
	Add(series common.Series, offset uint64, size uint64) error
	SortedID() []common.SeriesID
	// TODO: we need to record information like int/double, precision, min, max time etc. and we don't need to duplicate
	// series info like name tags for each entry, store in IndexEntries
	// TODO: use must means we will panic if use non exist ID
	MustEntries(id common.SeriesID) *IndexEntries
}

type IndexEntries struct {
	// TODO: series info,
	// TODO: we can't directly modify the tags returned by GetTags, because it is a reference
	tags    map[string]string
	entries []IndexEntry
}

type IndexEntry struct {
	// The absolute position where the data block starts
	Offset uint64
	// The size of the data block in bytes
	Size uint64
}

type LocalFileWriter struct {
	originalWriter io.WriteCloser
	w              io.Writer
	// block          BlockWriter
	index IndexWriter
}

type LocalFileBlockWriter struct {
}

type LocalFileIndexWriter struct {
	series map[common.SeriesID]*IndexEntries
}

func NewLocalFileWriter(w io.WriteCloser) *LocalFileWriter {
	return &LocalFileWriter{
		originalWriter: w,
		w:              bufio.NewWriter(w),
		index:          NewLocalFileIndexWriter(),
	}
}

func NewLocalFileIndexWriter() *LocalFileIndexWriter {
	return &LocalFileIndexWriter{
		series: map[common.SeriesID]*IndexEntries{},
	}
}

func (writer *LocalFileWriter) WriteSeries(series common.Series) error {
	// TODO: cast to different types
	return nil
}

func (writer *LocalFileWriter) Finalize() error {
	// TODO: write index to the end of the file
	return nil
}

func (writer *LocalFileWriter) Flush() error {
	// TODO: flush and sync underlying file
	return nil
}

func (writer *LocalFileWriter) Close() error {
	if err := writer.Flush(); err != nil {
		return errors.Wrap(err, "can't flush before close")
	}
	// TODO: close the underlying file
	return nil
}

func (idx *LocalFileIndexWriter) Add(series common.Series, offset uint64, size uint64) error {
	id := series.GetSeriesID()
	_, ok := idx.series[id]
	// create new IndexEntries if this series has not been added
	if !ok {

	}
	return nil
}

func (idx *LocalFileIndexWriter) SortedID() []common.SeriesID {
	keys := make([]common.SeriesID, 0, len(idx.series))
	for k := range idx.series {
		keys = append(keys, k)
	}
	sort.Sort(common.SeriesIDs(keys))
	return keys
}

func (idx *LocalFileIndexWriter) MustEntries(id common.SeriesID) *IndexEntries {
	entries, ok := idx.series[id]
	if !ok {
		log.Panicf("can't find entries for %d", id)
	}
	return entries
}
