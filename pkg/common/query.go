package common

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"github.com/pkg/errors"
)

type Filter struct {
	Type         string          `json:"type"`
	Key          string          `json:"key"`
	Value        string          `json:"value,omitempty"`
	Values       []string        `json:"values,omitempty"`
	LeftRaw      json.RawMessage `json:"l,omitempty"`
	RightRaw     json.RawMessage `json:"r,omitempty"`
	LeftOperand  *Filter         `json:"-"` // NOTE: must use pointer to avoid invalid recursive type Filter
	RightOperand *Filter         `json:"-"`
}

type filterAlias Filter

type Aggregator struct {
	Type   string `json:"type"`
	Window string `json:"window"` // TODO: change to time.Duration? or WindowRaw and Window with time.Duration
}

// Query is the query against single series if in `exact` mode, possible multiple series
// in `contains` mode
type Query struct {
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	MatchPolicy string            `json:"match_policy"`
	StartTime   int64             `json:"start_time,omitempty"`
	EndTime     int64             `json:"end_time,omitempty"`
	Filter      Filter            `json:"filter,omitempty"`
	Aggregator  Aggregator        `json:"aggregator,omitempty"`
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

func (filter *Filter) UnmarshalJSON(data []byte) error {
	// NOTE: need to use Alias like readRequest in `/server/service/read.go`, otherwise stackoverflow
	a := (*filterAlias)(filter)
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}
	if len(filter.LeftRaw) > 0 {
		var leftOperand Filter
		err := json.Unmarshal(filter.LeftRaw, &leftOperand)
		if err != nil {
			return errors.Wrap(err, "can't parse left operand")
		}
		filter.LeftOperand = &leftOperand
	}
	if len(filter.RightRaw) > 0 {
		var rightOperand Filter
		err := json.Unmarshal(filter.RightRaw, &rightOperand)
		if err != nil {
			return errors.Wrap(err, "can't parse right operand")
		}
		filter.RightOperand = &rightOperand
	}
	return nil
}
