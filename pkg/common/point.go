package common

import (
	"encoding/json"
)

// http://attilaolah.eu/2013/11/29/json-decoding-in-go/

type IntPoint struct {
	TimeNano int64
	V        int
}

// https://golang.org/pkg/encoding/json/#Marshaler
func (p *IntPoint) MarshalJSON() ([]byte, error) {
	// FIXME: I think there is way for not casting value to int64, maybe use sprinf
	return json.Marshal([2]int64{p.TimeNano, int64(p.V)})
	//return json.Marshal([2]json.Number{p.TimeNano, int64(p.V)})
}

// TODO: UseNumber seems only work for decoder https://golang.org/pkg/encoding/json/#Decoder.UseNumber

// TODO: unmarshaler does not return value?
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
	p.V = int(v)
	if err != nil {
		return err
	}
	return nil
}
