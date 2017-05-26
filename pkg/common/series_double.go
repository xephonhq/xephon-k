// Generated from series_int.go DO NOT EDIT!
package common

import "time"

// TODO: int series of other precision, maybe we should add millisecond to the default function as well
func NewDoubleSeries(name string) *DoubleSeries {
	return &DoubleSeries{
		SeriesMeta: SeriesMeta{
			Name:      name,
			Tags:      map[string]string{},
			Type:      TypeDoubleSeries,
			Precision: time.Millisecond.Nanoseconds(),
		},
	}
}
