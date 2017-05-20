package common

import (
	"encoding/json"
)

// IntPoint is a time, int pair but encoded as array in JSON format for space efficiency
// http://attilaolah.eu/2013/11/29/json-decoding-in-go/
type IntPoint struct {
	TimeNano int64
	V        int64
}

// MarshalJSON implements Marshaler interface
// https://golang.org/pkg/encoding/json/#Marshaler
func (p *IntPoint) MarshalJSON() ([]byte, error) {
	// FIXME: I think there is way for not casting value to int64, maybe use sprinf
	return json.Marshal([2]int64{p.TimeNano, p.V})
}

// UnmarshalJSON implements Unmarshaler interface
// https://golang.org/pkg/encoding/json/#Unmarshaler
func (p *IntPoint) UnmarshalJSON(data []byte) error {
	var tv [2]json.Number
	if err := json.Unmarshal(data, &tv); err != nil {
		return err
	}
	// FIXME: this error checking seems to be very low efficient
	t, err := tv[0].Int64()
	p.TimeNano = t
	if err != nil {
		return err
	}
	v, err := tv[1].Int64()
	p.V = v
	if err != nil {
		return err
	}
	return nil
}

// ByTime implements Sort interface for IntPoint
// https://golang.org/pkg/sort/
type ByTime []IntPoint

// Len implements Sort interface
func (p ByTime) Len() int {
	return len(p)
}

// Swap implements Sort interface
func (p ByTime) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less implements Sort interface
func (p ByTime) Less(i int, j int) bool {
	return p[i].TimeNano < p[j].TimeNano
}

type DoublePoint struct {
	T int64
	V float64
}
