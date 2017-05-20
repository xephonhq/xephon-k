package common

import (
	"testing"

	"encoding/json"

	asst "github.com/stretchr/testify/assert"
	"time"
)

func TestIntSeries_JSON(t *testing.T) {
	assert := asst.New(t)

	s := IntSeries{
		Name:   "cpi",
		Tags:   map[string]string{"os": "ubuntu"},
		Points: []IntPoint{{TimeNano: 1359788400000, V: 1}, {TimeNano: 1359788400001, V: 2}},
	}
	j, err := json.Marshal(s)
	assert.Nil(err)
	assert.Equal(`{"name":"cpi","tags":{"os":"ubuntu"},"points":[[1359788400000,1],[1359788400001,2]]}`, string(j))

	s2 := IntSeries{}
	err = json.Unmarshal(j, &s2)
	assert.Nil(err)
	assert.Equal("ubuntu", s2.Tags["os"])
	assert.Equal(int64(1359788400000), s2.Points[0].TimeNano)
}

func TestIntSeries_Hash(t *testing.T) {
	assert := asst.New(t)

	tags := make(map[string]string)
	tags["os"] = "ubuntu"
	tags["region"] = "us-east"
	p1 := IntPoint{TimeNano: 1359788400000, V: 1}
	p2 := IntPoint{TimeNano: 1359788400001, V: 2}
	ps1 := []IntPoint{p1}
	ps2 := []IntPoint{p2}
	s1 := IntSeries{Name: "cpi", Tags: tags, Points: ps1}
	s2 := IntSeries{Name: "cpi", Tags: tags, Points: ps2}
	s3 := IntSeries{Name: "ipc", Tags: tags, Points: ps2} // different name
	tags2 := make(map[string]string)
	s4 := IntSeries{Name: "cpi", Tags: tags2, Points: ps1} // different tag
	tags3 := make(map[string]string)
	tags3["os"] = "ubuntu"
	tags3["region"] = "us-east"
	// same tag, different tag object, and simply range on tags should have different result on every run, that's why we
	// sort it before we calculate the hash
	s5 := IntSeries{Name: "cpi", Tags: tags3, Points: ps1}
	assert.Equal(s1.Hash(), s2.Hash())
	assert.NotEqual(s1.Hash(), s3.Hash())
	assert.NotEqual(s1.Hash(), s4.Hash())
	assert.Equal(s1.Hash(), s5.Hash())
}

func TestDoubleSeries_JSON(t *testing.T) {
	assert := asst.New(t)
	s := DoubleSeries{
		Name:      "cpi",
		Tags:      map[string]string{"os": "ubuntu"},
		Precision: time.Millisecond,
		Points:    []DoublePoint{{T: 1359788400000, V: 1.0}, {T: 1359788400001, V: 2.08}},
	}
	j, err := json.Marshal(s)
	assert.Nil(err)
	assert.Equal(`{"name":"cpi","tags":{"os":"ubuntu"},"precision":1000000,"points":[[1359788400000,1.000000],[1359788400001,2.080000]]}`,
		string(j))

	s2 := DoubleSeries{}
	err = json.Unmarshal(j, &s2)
	assert.Nil(err)
	assert.Equal(s.Tags["os"], s2.Tags["os"])
	assert.Equal(s.Points[0].T, s2.Points[0].T)
}
