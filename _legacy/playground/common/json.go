package common

import (
	"fmt"
	"encoding/json"
)

func (m *IntPoint) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[%d,%d]", m.T, m.V)), nil
}

func (m *IntPoint) UnmarshalJSON(data []byte) error {
	var tv [2]json.Number
	if err := json.Unmarshal(data, &tv); err != nil {
		return err
	}
	// FIXME: this error checking seems to be very low efficient
	t, err := tv[0].Int64()
	m.T = t
	if err != nil {
		return err
	}
	v, err := tv[1].Int64()
	m.V = v
	if err != nil {
		return err
	}
	return nil
}
