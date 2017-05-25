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

func TestMeta_Marshal(t *testing.T) {
	m := Meta{
		Id:   123,
		Type: 1,
		Name: "haha",
		Tags: map[string]string{"os": "ubuntu"},
	}
	b, err := m.Marshal()
	if err != nil {
		t.Fatal("can't marshal to bytes")
	}
	m2 := Meta{}
	err = m2.Unmarshal(b)
	if err != nil {
		t.Fatal("can't unmarshal from bytes")
	}
	if m.Id != m2.Id {
		t.Fatal("should match")
	}
}

func TestIntSeries_Marshal(t *testing.T) {
	intS := IntSeries{
		Meta: Meta{
			Id:   123,
			Type: 1,
			Name: "haha",
			Tags: map[string]string{"os": "ubuntu"},
		},
		Points: []IntPoint{{T: 1, V: 1}, {T: 1, V: 2}},
	}
	b, err := intS.Marshal()
	if err != nil {
		t.Fatal("can't marshal int series")
	}
	intS2 := IntSeries{}
	err = intS2.Unmarshal(b)
	if err != nil {
		t.Fatal("can't unmarshal int series")
	}
	if intS.Meta.Id != intS2.Meta.Id {
		t.Fatal("should match")
	}
}
