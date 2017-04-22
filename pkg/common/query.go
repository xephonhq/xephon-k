package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
)

// Query is the query against single series if in `exact` mode, possible multiple series
// in `contains` mode
type Query struct {
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	MatchPolicy string            `json:"match_policy"`
	StartTime   int64             `json:"start_time,omitempty"`
	EndTime     int64             `json:"end_time,omitempty"`
}

// QueryResult contains the original query and number of series matched
type QueryResult struct {
	Query
	Matched int `json:"matched"`
}

// Hash return the same result as IntSeries's hash function
func (query *Query) Hash() SeriesID {
	// TODO: this is copied from series Hash
	h := md5.New()
	io.WriteString(h, query.Name)
	keys := make([]string, len(query.Tags))
	i := 0
	for k := range query.Tags {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		io.WriteString(h, k)
		io.WriteString(h, query.Tags[k])
	}
	return SeriesID(fmt.Sprintf("%x", h.Sum(nil)))
}
