package common

import (
	"testing"

	"encoding/json"

	asst "github.com/stretchr/testify/assert"
	"time"
)

// FIXED: tags are passed by reference not value https://github.com/xephonhq/xephon-k/issues/40
func TestIntSeries_GetTags(t *testing.T) {
	assert := asst.New(t)

	s := IntSeries{
		SeriesMeta: SeriesMeta{
			Name: "cpi",
			Tags: map[string]string{"os": "ubuntu"},
		},
		Points: []IntPoint{{T: 1359788400000, V: 1}, {T: 1359788400001, V: 2}},
	}
	// NOTE: tagsCopy and s.Tags refer to the same underlying data
	tagsCopy := s.GetTags()
	tagsCopy["os"] = "fedora"
	assert.Equal("fedora", s.Tags["os"])
}

func TestIntSeries_GetSeriesID(t *testing.T) {
	assert := asst.New(t)

	s := IntSeries{}
	err := json.Unmarshal([]byte(`{"name":"cpi","tags":{"os":"ubuntu"},"points":[[1359788400000,1],[1359788400001,2]]}`), &s)
	assert.Nil(err)

	assert.Equal(Hash(&s), s.GetSeriesID())
}

func TestIntSeries_JSON(t *testing.T) {
	assert := asst.New(t)

	s := IntSeries{
		SeriesMeta: SeriesMeta{
			Name: "cpi",
			Tags: map[string]string{"os": "ubuntu"},
		},
		Points: []IntPoint{{T: 1359788400000, V: 1}, {T: 1359788400001, V: 2}},
	}
	j, err := json.Marshal(s)
	assert.Nil(err)
	assert.Equal(`{"meta":{"name":"cpi","tags":{"os":"ubuntu"}},"points":[[1359788400000,1],[1359788400001,2]]}`, string(j))

	s2 := IntSeries{}
	err = json.Unmarshal(j, &s2)
	assert.Nil(err)
	assert.Equal("ubuntu", s2.Tags["os"])
	assert.Equal(int64(1359788400000), s2.Points[0].T)
}

func TestIntSeries_Hash(t *testing.T) {
	assert := asst.New(t)
	// TODO: change to a table test
	tags := map[string]string{"os": "ubuntu", "region": "us-east"}
	p1 := IntPoint{T: 1359788400000, V: 1}
	p2 := IntPoint{T: 1359788400001, V: 2}
	s1 := IntSeries{SeriesMeta: SeriesMeta{Name: "cpi", Tags: tags}, Points: []IntPoint{p1}}
	s2 := IntSeries{SeriesMeta: SeriesMeta{Name: "cpi", Tags: tags}, Points: []IntPoint{p2}}
	s3 := IntSeries{SeriesMeta: SeriesMeta{Name: "ipc", Tags: tags}, Points: []IntPoint{p2}} // different name
	tags2 := make(map[string]string)
	s4 := IntSeries{SeriesMeta: SeriesMeta{Name: "cpi", Tags: tags2}, Points: []IntPoint{p1}} // different tag
	tags3 := map[string]string{"os": "ubuntu", "region": "us-east"}
	// same tag, different tag object, and simply range on tags should have different result on every run,
	// that's why we sort it before we calculate the hash
	s5 := IntSeries{SeriesMeta: SeriesMeta{Name: "cpi", Tags: tags3}, Points: []IntPoint{p1}}
	assert.Equal(Hash(&s1), Hash(&s2))
	assert.NotEqual(Hash(&s1), Hash(&s3))
	assert.NotEqual(Hash(&s1), Hash(&s4))
	assert.Equal(Hash(&s1), Hash(&s5))
}

func TestDoubleSeries_JSON(t *testing.T) {
	assert := asst.New(t)
	s := DoubleSeries{
		SeriesMeta: SeriesMeta{
			Name:      "cpi",
			Tags:      map[string]string{"os": "ubuntu"},
			Precision: time.Millisecond.Nanoseconds(),
		},
		Points: []DoublePoint{{T: 1359788400000, V: 1.0}, {T: 1359788400001, V: 2.08}},
	}
	j, err := json.Marshal(s)
	assert.Nil(err)
	assert.Equal(`{"meta":{"precision":1000000,"name":"cpi","tags":{"os":"ubuntu"}},"points":[[1359788400000,1.000000],[1359788400001,2.080000]]}`,
		string(j))

	s2 := DoubleSeries{}
	err = json.Unmarshal(j, &s2)
	assert.Nil(err)
	assert.Equal(s.Tags["os"], s2.Tags["os"])
	assert.Equal(s.Points[0].T, s2.Points[0].T)
}
