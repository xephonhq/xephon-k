package common

import (
	"encoding/json"
	"fmt"
	"time"
)

// check interface
var _ Series = (*SeriesMeta)(nil)
var _ Series = (*RawSeries)(nil)
var _ Series = (*IntSeries)(nil)
var _ Series = (*DoubleSeries)(nil)

const (
	_ = iota
	TypeIntSeries
	TypeDoubleSeries
	TypeBoolSeries
	TypeStringSeries
)

type Series interface {
	Hashable
	GetSeriesType() int
	// NOTE: series decoded from JSON has 0 as SeriesID, so the implementation would recalculate the Hash
	GetSeriesID() SeriesID
}

type SeriesMeta struct {
	id         SeriesID
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type"`
	Precision  time.Duration     `json:"precision"`
}

type RawSeries struct {
	SeriesMeta
	Points json.RawMessage `json:"points"`
}

type IntSeries struct {
	SeriesMeta
	Points []IntPoint `json:"points"`
}

type DoubleSeries struct {
	SeriesMeta
	Points []DoublePoint `json:"points"`
}

// TODO: int series of other precision, maybe we should add millisecond to the default function as well
func NewIntSeries(name string) *IntSeries {
	// return &IntSeries{
	// 	Name:       name,
	// 	Tags:       make(map[string]string, 1),
	// 	SeriesType: TypeIntSeries,
	// 	Precision:  time.Millisecond,
	// }

	s := IntSeries{}
	s.Name = name
	s.Tags = map[string]string{}
	s.SeriesType = TypeIntSeries
	s.Precision = time.Millisecond
	return &s
}

func NewDoubleSeries(name string) *DoubleSeries {
	// return &DoubleSeries{
	// 	Name:       name,
	// 	Tags:       make(map[string]string, 1),
	// 	SeriesType: TypeDoubleSeries,
	// 	Precision:  time.Millisecond,
	// }
	s := DoubleSeries{}
	s.Name = name
	s.Tags = map[string]string{}
	s.SeriesType = TypeDoubleSeries
	s.Precision = time.Millisecond
	return &s
}

func (series *SeriesMeta) GetName() string {
	return series.Name
}

func (series *SeriesMeta) GetTags() map[string]string {
	return series.Tags
}

func (series *SeriesMeta) GetSeriesType() int {
	return series.SeriesType
}

func (series *SeriesMeta) GetSeriesID() SeriesID {
	if series.id == 0 {
		series.id = Hash(series)
	}
	return series.id
}

func SeriesTypeString(seriesType int) string {
	switch seriesType {
	case TypeIntSeries:
		return "int"
	case TypeDoubleSeries:
		return "double"
	case TypeBoolSeries:
		return "bool"
	case TypeStringSeries:
		return "string"
	default:
		return fmt.Sprintf("unknown: %d", seriesType)
	}
}
