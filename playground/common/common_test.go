package common

import (
	"testing"
	"encoding/json"
)

func TestIntPoint_Marshal(t *testing.T) {
	p := IntPoint{T: 123, V: 1}
	b, err := p.Marshal()
	if err != nil {
		t.Fatal("cant marshal to bytes")
	}
	p2 := IntPoint{}
	err = p2.Unmarshal(b)
	if err != nil {
		t.Fatal("can't unmarshal from bytes")
	}
	// the generated type can co-exists with custom JSON
	bj, err := json.Marshal(p)
	t.Log(string(bj))
}
