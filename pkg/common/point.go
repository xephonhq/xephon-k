package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// IntPoint is a time, int pair. Its time precision is based on the series it belows to.
// It is encoded as array in JSON format for space efficiency.
type IntPoint struct {
	T int64
	V int64
}

// DoublePoint is a time, double pair. Used the same way as IntPoint.
type DoublePoint struct {
	T int64
	V float64
}

// MarshalJSON implements Marshaler interface. Pointed is encoded as number array. i.e. [1359788400000,1]
// https://golang.org/pkg/encoding/json/#Marshaler
func (p *IntPoint) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[%d,%d]", p.T, p.V)), nil
}

func (p *IntPoint) MarshalJSON2() ([]byte, error) {
	return json.Marshal([2]int64{p.T, p.V})
}

func (p *IntPoint) MarshalJSON3() ([]byte, error) {
	return []byte("[" + strconv.FormatInt(p.T, 10) + "," + strconv.FormatInt(p.V, 10) + "]"), nil
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
	p.T = t
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

// IntPoints implements sort.Interface for IntPoint
// https://golang.org/pkg/sort/
type IntPoints []IntPoint

// Len implements Sort interface
func (p IntPoints) Len() int {
	return len(p)
}

// Swap implements Sort interface
func (p IntPoints) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less implements Sort interface
func (p IntPoints) Less(i int, j int) bool {
	return p[i].T < p[j].T
}

// MarshalJSON implements Marshaler interface. Point is encoded as number array. i.e. []
func (p *DoublePoint) MarshalJSON() ([]byte, error) {
	// TODO: precision of double value, need to copy the code in `json/encode.go`
	return []byte(fmt.Sprintf("[%d,%f]", p.T, p.V)), nil
}

// UnmarshalJSON implements Unmarshaler interface.
func (p *DoublePoint) UnmarshalJSON(data []byte) error {
	var tv [2]json.Number
	if err := json.Unmarshal(data, &tv); err != nil {
		return err
	}
	t, err := tv[0].Int64()
	if err != nil {
		return err
	}
	p.T = t
	v, err := tv[1].Float64()
	if err != nil {
		return err
	}
	p.V = v
	return nil
}

// DoublePoints implements sort.Interface for DoublePoint
type DoublePoints []DoublePoint

func (p DoublePoints) Len() int {
	return len(p)
}

func (p DoublePoints) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p DoublePoints) Less(i int, j int) bool {
	return p[i].T < p[j].T
}
