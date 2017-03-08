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
