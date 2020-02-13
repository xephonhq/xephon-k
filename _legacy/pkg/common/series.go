package common

import (
	"encoding/json"
	"fmt"
)

// check interface
// var _ Series = (*SeriesMeta)(nil)
var _ Series = (*RawSeries)(nil)
var _ Series = (*IntSeries)(nil)
var _ Series = (*DoubleSeries)(nil)
var _ Series = (*IntSeriesColumnar)(nil)
var _ Series = (*DoubleSeriesColumnar)(nil)

const (
	TypeRawSeries = iota
	TypeIntSeries
	TypeDoubleSeries
	TypeBoolSeries
	TypeStringSeries
)

type Series interface {
	Hashable
	GetSeriesType() int64
	GetSeriesID() SeriesID // NOTE: series decoded from JSON has 0 as SeriesID, so the implementation would recalculate the Hash
	GetMetaCopy() SeriesMeta
	GetMinTime() int64
	GetMaxTime() int64
	Length() int
	PrintPoints()
}

type SeriesColumnar interface {
	Series
	IsColumnar() bool
}

// NOTE: The following struct for series are generated by proto in common.pb.go
//  SeriesMeta
//  IntSeries
//  DoubleSeries
//  IntSeriesColumnar
//	DoubleSeriesColumnar

// RawSeries is an intermediate struct for decoding json with mixed type of series in an array
type RawSeries struct {
	SeriesMeta `json:"meta"`
	Points     json.RawMessage `json:"points"`
}

func SeriesTypeString(seriesType int64) string {
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

func (m *SeriesMeta) GetName() string {
	return m.Name
}

func (m *SeriesMeta) GetTags() map[string]string {
	return m.Tags
}

func (m *SeriesMeta) GetSeriesType() int64 {
	return m.Type
}

func (m *SeriesMeta) GetSeriesID() SeriesID {
	if m.Id == 0 {
		m.Id = uint64(Hash(m))
	}
	return SeriesID(m.Id)
}

func (m *SeriesMeta) GetMetaCopy() SeriesMeta {
	newMap := make(map[string]string, len(m.Tags))
	for k, v := range m.Tags {
		newMap[k] = v
	}
	return SeriesMeta{
		Id:        m.Id,
		Type:      m.Type,
		Precision: m.Precision,
		Name:      m.Name, // TODO: string is actually a reference to underlying slice, I don't know if is there is impact on GC if we don't make a deep copy
		Tags:      newMap,
	}
}

func (m *RawSeries) GetMinTime() int64 {
	return 0
}

func (m *RawSeries) GetMaxTime() int64 {
	return 0
}

func (m *RawSeries) Length() int {
	return 0
}

func (m *RawSeries) PrintPoints() {
	fmt.Print("FIXME: shoud not see PrintPoints called on RawSeries")
	fmt.Print(m.Points)
}