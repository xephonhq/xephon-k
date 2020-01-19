package common

import (
	"fmt"
	"time"
)

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

func (m *IntSeries) Length() int {
	return len(m.Points)
}

func (m *IntSeries) PrintPoints() {
	for i, p := range m.Points {
		fmt.Printf("%d, %d, %v\n", i, p.T, p.V)
	}
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

func (m *IntSeriesColumnar) IsColumnar() bool {
	return true
}

func (m *IntSeriesColumnar) Length() int {
	return len(m.T)
}

func (m *IntSeriesColumnar) PrintPoints() {
	for i := 0; i < len(m.T); i++ {
		fmt.Printf("%d, %d, %v\n", i, m.T[i], m.V[i])
	}
}
