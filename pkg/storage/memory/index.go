package memory

import (
	"github.com/xephonhq/xephon-k/pkg/common"
)

// Index is a map of inverted index with tag name as key and tag value as term for the inverted index
type Index struct {
	tagKeyIndex     map[string]map[string]bool // map[string]bool is used as set
	invertedIndexes map[string]*InvertedIndex
}

// InvertedIndex use Term for tag value postings for a list of sorted series ID
// TODO: Series ID should use locality sensitive hashing https://en.wikipedia.org/wiki/Locality-sensitive_hashing
type InvertedIndex struct {
	Term     string
	Postings []common.SeriesID
}

var initialPostingSize = 10

func NewIndex(capacity int) *Index {
	return &Index{
		tagKeyIndex:     make(map[string]map[string]bool, capacity),
		invertedIndexes: make(map[string]*InvertedIndex, capacity),
	}
}

func newInvertedIndex(term string) *InvertedIndex {
	return &InvertedIndex{
		Term:     term,
		Postings: make([]common.SeriesID, 0, initialPostingSize),
	}
}

func (idx *Index) Add(id common.SeriesID, tagKey string, tagValue string) {
	// update tagKeyIndex
	_, ok := idx.tagKeyIndex[tagKey]
	if !ok {
		idx.tagKeyIndex[tagKey] = make(map[string]bool)
	}
	idx.tagKeyIndex[tagKey][tagValue] = true

	// TODO: should add separator, in Prometheus `db.go` it's `const sep = '\xff'`
	term := tagKey + tagValue
	// create the inverted index if not exists
	_, ok = idx.invertedIndexes[term]
	if !ok {
		idx.invertedIndexes[term] = newInvertedIndex(term)
	}
	idx.invertedIndexes[term].Add(id)
}

// TODO: actually we can have a fixed size map to cache the hot series, so there is no need to lookup if the id is already in there
func (iidx *InvertedIndex) Add(id common.SeriesID) {
	// binary search and insert the value if not found
	low, high := 0, len(iidx.Postings)
	for low < high {
		// TODO: use custom compare function or compare it directly using <
		mid := low + (high-low)/2 // avoid overflow, copied from `src/sort/search.go` sort.Search
		if iidx.Postings[mid] >= id {
			high = mid
		} else {
			low = mid + 1
		}
	}

	// not found
	if low == len(iidx.Postings) {
		iidx.Postings = append(iidx.Postings, id)
		return
	} else if iidx.Postings[low] != id {
		// insert it to the slice https://github.com/golang/go/wiki/SliceTricks#insert
		iidx.Postings = append(iidx.Postings, id) // we append id here, but any value is ok, it will be overwritten by following copy
		copy(iidx.Postings[low+1:], iidx.Postings[low:])
		iidx.Postings[low] = id

	}

	// found
	// TODO: should have some sort of cache
	return
}
