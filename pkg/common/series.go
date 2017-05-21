package common

import (
	"sort"
	"time"
)

// check interface
var _ Series = (*IntSeries)(nil)
var _ Series = (*DoubleSeries)(nil)

type Series interface {
	Hashable
}

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

func NewDoubleSeries(name string) *DoubleSeries {
	return &DoubleSeries{
		Name:      name,
		Tags:      make(map[string]string, 1),
		Precision: time.Millisecond,
	}
}

func (series *IntSeries) GetName() string {
	return series.Name
}

func (series *IntSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *DoubleSeries) GetName() string {
	return series.Name
}

func (series *DoubleSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *DoubleSeries) Hash() SeriesID {
	// TODO: copied from double series
	h := NewInlineFNV64a()
	h.Write([]byte(series.Name))
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
		h.Write([]byte(k))
		h.Write([]byte(series.Tags[k]))
	}
	return SeriesID(h.Sum64())
}
