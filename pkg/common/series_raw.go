// Generated from series_int.go DO NOT EDIT!
// NOTE: we have edited the min, max time method
package common

import "time"

// TODO: int series of other precision, maybe we should add millisecond to the default function as well
func NewRawSeries(name string) *RawSeries {
	return &RawSeries{
		SeriesMeta: SeriesMeta{
			Name:      name,
			Tags:      map[string]string{},
			Type:      TypeRawSeries,
			Precision: time.Millisecond.Nanoseconds(),
		},
	}
}

func (m *RawSeries) GetMinTime() int64 {
	return 0
}

func (m *RawSeries) GetMaxTime() int64 {
	return 0
}
