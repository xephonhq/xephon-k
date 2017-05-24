package common

import (
	"encoding/json"
	"fmt"
	"time"
)

// check interface
var _ Series = (*MetaSeries)(nil)
var _ Series = (*IntSeries)(nil)
var _ Series = (*DoubleSeries)(nil)

const (
	_                = iota
	TypeIntSeries
	TypeDoubleSeries
	TypeBoolSeries
	TypeStringSeries
)

// check interface
var _ Series = (*MetaSeries)(nil)
var _ Series = (*IntSeries)(nil)
var _ Series = (*DoubleSeries)(nil)

type Series interface {
	Hashable
	GetSeriesType() int
	// TODO: replace hash with GetSeriesID and see if it works with series decoded from JSON
	GetSeriesID() SeriesID
}

type MetaSeries struct {
	id         SeriesID
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type"`
	Precision  time.Duration     `json:"precision"`
	Points     json.RawMessage   `json:"points"`
}

type IntSeries struct {
	id         SeriesID
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type,omitempty"`
	Precision  time.Duration     `json:"precision,omitempty"`
	Points     []IntPoint        `json:"points"`
}

type DoubleSeries struct {
	id         SeriesID
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type,omitempty"`
	Precision  time.Duration     `json:"precision,omitempty"`
	Points     []DoublePoint     `json:"points"`
}

// TODO: int series of other precision, maybe we should add millisecond to the default function as well
func NewIntSeries(name string) *IntSeries {
	return &IntSeries{
		Name:       name,
		Tags:       make(map[string]string, 1),
		SeriesType: TypeIntSeries,
		Precision:  time.Millisecond,
	}
}

func NewDoubleSeries(name string) *DoubleSeries {
	return &DoubleSeries{
		Name:       name,
		Tags:       make(map[string]string, 1),
		SeriesType: TypeDoubleSeries,
		Precision:  time.Millisecond,
	}
}

func (series *MetaSeries) GetName() string {
	return series.Name
}

func (series *MetaSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *MetaSeries) GetSeriesType() int {
	return series.SeriesType
}

func (series *MetaSeries) GetSeriesID() SeriesID {
	if series.id == 0 {
		series.id = Hash(series)
	}
	return series.id
}

func (series *IntSeries) GetName() string {
	return series.Name
}

func (series *IntSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *IntSeries) GetSeriesType() int {
	// TODO: do we still need the variable if we return constant, and should we check consistency for that
	return TypeIntSeries
}

func (series *IntSeries) GetSeriesID() SeriesID {
	if series.id == 0 {
		series.id = Hash(series)
	}
	return series.id
}

func (series *DoubleSeries) GetName() string {
	return series.Name
}

func (series *DoubleSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *DoubleSeries) GetSeriesType() int {
	return TypeDoubleSeries
}

func (series *DoubleSeries) GetSeriesID() SeriesID {
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
