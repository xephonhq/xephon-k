package disk

import (
	"github.com/xephonhq/xephon-k/pkg/common"
)

type IndexEntries struct {
	// TODO: index entries may have the overall aggregation of all the index entries, like min, max, time
	// TODO: use protobuf to encode
	meta    common.SeriesMeta
	entries []IndexEntry
}

type IndexEntry struct {
	// TODO: aggregated data
	// Offset is the absolute position where the data block starts
	Offset uint64
	// Size is the size of the data block in bytes
	Size uint64
}
