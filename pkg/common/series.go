package common

import (
	"encoding/json"
	"time"
)

// check interface
var _ Series = (*MetaSeries)(nil)
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
}

type MetaSeries struct {
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type"`
	Precision  time.Duration     `json:"precision"`
	Points     json.RawMessage   `json:"points"`
}

type IntSeries struct {
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	SeriesType int               `json:"type,omitempty"`
	Precision  time.Duration     `json:"precision,omitempty"`
	Points     []IntPoint        `json:"points"`
}

type DoubleSeries struct {
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

func (series *DoubleSeries) GetName() string {
	return series.Name
}

func (series *DoubleSeries) GetTags() map[string]string {
	return series.Tags
}

func (series *DoubleSeries) GetSeriesType() int {
	return TypeDoubleSeries
}
