package memory

// Index is a map of inverted index with tag name as key and tag value as term for the inverted index
type Index map[string]InvertedIndex

// InvertedIndex use Term for tag value postings for a list of sorted series ID
// TODO: Series ID should use locality hashsing
type InvertedIndex struct {
	Term     string
	Postings []SeriesID
}
