package common

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestIntSeries_JSON(t *testing.T) {
	asst := assert.New(t)

	tags := make(map[string]string)
	tags["os"] = "ubuntu"
	p1 := IntPoint{TimeNano: 1359788400000, V: 1}
	p2 := IntPoint{TimeNano: 1359788400001, V: 2}
	points := []IntPoint{p1, p2}
	s := IntSeries{Name: "cpi", Tags: tags, Points: points}
	j, err := json.Marshal(s)
	asst.Nil(err)
	asst.Equal(`{"name":"cpi","tags":{"os":"ubuntu"},"points":[[1359788400000,1],[1359788400001,2]]}`, string(j))

	s2 := IntSeries{}
	err = json.Unmarshal(j, &s2)
	asst.Nil(err)
	asst.Equal("ubuntu", s2.Tags["os"])
	asst.Equal(int64(1359788400000), s2.Points[0].TimeNano)
}

func TestIntSeries_Hash(t *testing.T) {
	asst := assert.New(t)

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
	asst.Equal(s1.Hash(), s2.Hash())
	asst.NotEqual(s1.Hash(), s3.Hash())
	asst.NotEqual(s1.Hash(), s4.Hash())
	asst.Equal(s1.Hash(), s5.Hash())

}
