package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
)

type IntSeries struct {
	Name   string            `json:"name"`
	Tags   map[string]string `json:"tags"`
	Points []IntPoint        `json:"points"`
}

type DoubleSeries struct {
	Name   string            `json:"name"`
	Tags   map[string]string `json:"tags"`
	Points []DoublePoint     `json:"points"`
}

// Hash returns one result for series have same name and tags
func (series *IntSeries) Hash() string {
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
	return fmt.Sprintf("%x", h.Sum(nil))
}
