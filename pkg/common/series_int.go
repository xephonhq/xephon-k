package common

import "time"

// TODO: int series of other precision, maybe we should add millisecond to the default function as well
func NewIntSeries(name string) *IntSeries {
	return &IntSeries{
		SeriesMeta: SeriesMeta{
			Name:      name,
			Tags:      map[string]string{},
			Type:      TypeIntSeries,
			Precision: time.Millisecond.Nanoseconds(),
		},
	}
}

func (m *IntSeries) GetMinTime() int64 {
	if len(m.Points) == 0 {
		return 0
	}
	return m.Points[0].T
}

func (m *IntSeries) GetMaxTime() int64 {
	if len(m.Points) == 0 {
		return 0
	}
	return m.Points[len(m.Points)-1].T
}

func (m *IntSeriesColumnar) GetMinTime() int64 {
	if len(m.T) == 0 {
		return 0
	}
	return m.T[0]
}

func (m *IntSeriesColumnar) GetMaxTime() int64 {
	if len(m.T) == 0 {
		return 0
	}
	return m.T[len(m.T)-1]
}
