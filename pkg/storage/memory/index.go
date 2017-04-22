package memory

import "github.com/xephonhq/xephon-k/pkg/common"

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

func (idx *Index) Add(id common.SeriesID, tagName string, tagValue string) {
	// TODO: should add separator, in Prometheus `db.go` it's `const sep = '\xff'`
	term := tagName + tagValue
	// create the inverted index if not exists
	iidx, ok := idx.invertedIndexes[term]
	if !ok {
		iidx = newInvertedIndex(term)
	}
	iidx.Add(id, term)
	// TODO: update tagKeyIndex
}

// TODO: actually we can have a fixed size map to cache the hot series, so there is no need to lookup if the id is already in there
func (iidx *InvertedIndex) Add(id common.SeriesID, term string) {
	// TODO: add the id and keep the list sorted
}
