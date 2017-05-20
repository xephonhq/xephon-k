package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"time"
)

type IntSeries struct {
	Name   string            `json:"name"`
	Tags   map[string]string `json:"tags"`
	Points []IntPoint        `json:"points"`
}

type DoubleSeries struct {
	Name      string            `json:"name"`
	Tags      map[string]string `json:"tags"`
	Precision time.Duration     `json:"precision"`
	Points    []DoublePoint     `json:"points"`
}

// TODO: int series of other precision, maybe we should add millisecond to the default function as well
func NewIntSeries(name string) *IntSeries {
	return &IntSeries{
		Name: name,
		Tags: make(map[string]string, 1),
	}
}

// Hash returns one result for series have same name and tags
func (series *IntSeries) Hash() SeriesID {
	// TODO: more efficient way for hashing, every time we hash, we sort it, and using []byte
	// should be more efficient than string
	h := md5.New()
	io.WriteString(h, series.Name)
	keys := make([]string, len(series.Tags))
	i := 0
	// NOTE: use range on map has different order of keys on every run, except you only have one key,
	// thus we need to sort the keys when we calculate the hash
	for k := range series.Tags {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		io.WriteString(h, k)
		io.WriteString(h, series.Tags[k])
	}
	return SeriesID(fmt.Sprintf("%x", h.Sum(nil)))
}
