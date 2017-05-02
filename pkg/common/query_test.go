package common

import (
	"testing"

	"encoding/json"

	asst "github.com/stretchr/testify/assert"
)

func TestQuery_JSON(t *testing.T) {
	assert := asst.New(t)
	tags := make(map[string]string)
	tags["os"] = "ubuntu"
	q := Query{Tags: tags, Name: "cpu.idle", MatchPolicy: "exact"}
	j, err := json.Marshal(q)
	assert.Nil(err)
	assert.Equal(`{"name":"cpu.idle","tags":{"os":"ubuntu"},"match_policy":"exact"}`, string(j))
	qr := QueryResult{Query: q, Matched: 1}
	j, err = json.Marshal(qr)
	assert.Nil(err)
	assert.Equal(`{"name":"cpu.idle","tags":{"os":"ubuntu"},"match_policy":"exact","matched":1}`, string(j))

	// unmarshal an array of queries
	queriesData := `[
			{
				"name":"cpi",
				"tags":{"machine":"machine-01","os":"ubuntu"},
				"match_policy": "exact",
				"start_time": 1493363958000,
				"end_time": 1494363958000
			},
			{
				"name":"cpi",
				"tags":{"machine":"machine-02","os":"ubuntu"},
				"match_policy": "exact",
				"start_time": 1493363958000,
				"end_time": 1494363958000
			}
	]`
	var queries []Query
	err = json.Unmarshal([]byte(queriesData), &queries)
	assert.Nil(err)
	assert.Equal(2, len(queries))
}

func TestQuery_Hash(t *testing.T) {
	assert := asst.New(t)

	tags1 := make(map[string]string)
	tags1["os"] = "ubuntu"
	tags1["region"] = "us"
	tags2 := make(map[string]string)
	tags2["os"] = "ubuntu"
	tags2["region"] = "us"

	q := Query{Tags: tags1, Name: "cpu.idle", MatchPolicy: "exact"}
	s := IntSeries{Tags: tags2, Name: "cpu.idle"}
	assert.Equal(q.Hash(), s.Hash())
}
