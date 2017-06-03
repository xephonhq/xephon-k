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

func (m *DoubleSeries) GetMinTime() int64 {
	if len(m.Points) == 0 {
		return 0
	}
	return m.Points[0].T
}

func (m *DoubleSeries) GetMaxTime() int64 {
	if len(m.Points) == 0 {
		return 0
	}
	return m.Points[len(m.Points)-1].T
}

func (m *DoubleSeriesColumnar) GetMinTime() int64 {
	if len(m.T) == 0 {
		return 0
	}
	return m.T[0]
}

func (m *DoubleSeriesColumnar) GetMaxTime() int64 {
	if len(m.T) == 0 {
		return 0
	}
	return m.T[len(m.T)-1]
}
