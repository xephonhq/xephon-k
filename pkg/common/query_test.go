package common

import (
	"testing"

	"encoding/json"

	asst "github.com/stretchr/testify/assert"
)

func TestFilter_UnmarshalJSON(t *testing.T) {
	assert := asst.New(t)
	filterData := `{
		"type": "and",
		"l": {
			"type": "tag_match",
			"key": "region",
			"value": "en-us"
		},
		"r": {
			"type": "or",
			"l": {
				"type": "tag_match",
				"key": "connected_to",
				"value": "en-us"
			},
			"r": {
				"type": "in",
				"key": "machine_type",
				"values": ["switch", "router"]
			}
		}
	}`
	var filter Filter
	err := json.Unmarshal([]byte(filterData), &filter)
	assert.Nil(err)
	assert.Equal("or", filter.RightOperand.Type)
	assert.Equal("tag_match", filter.RightOperand.LeftOperand.Type)
	assert.Equal(2, len(filter.RightOperand.RightOperand.Values))
}

func TestQuery_JSON(t *testing.T) {
	assert := asst.New(t)
	tags := make(map[string]string)
	tags["os"] = "ubuntu"
	q := Query{Tags: tags, Name: "cpu.idle", MatchPolicy: "exact"}
	_, err := json.Marshal(q)
	assert.Nil(err)
	// FIXME: there are empty filter and aggregator in result json, omitempty does not work for custom type
	//assert.Equal(`{"name":"cpu.idle","tags":{"os":"ubuntu"},"match_policy":"exact"}`, string(j))
	qr := QueryResult{Query: q, Matched: 1}
	_, err = json.Marshal(qr)
	assert.Nil(err)
	//assert.Equal(`{"name":"cpu.idle","tags":{"os":"ubuntu"},"match_policy":"exact","matched":1}`, string(j))

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
				"name": "cpu.usage",
				"match_policy": "filter",
				"filter": {
					"type": "and",
					"l": {
						"type": "tag_match",
						"key": "region",
						"value": "en-us"
					},
					"r": {
						"type": "or",
						"l": {
							"type": "tag_match",
							"key": "connected_to",
							"value": "en-us"
						},
						"r": {
							"type": "in",
							"key": "machine_type",
							"values": ["switch", "router"]
						}
					}
				},
				"aggregator": {
					"type": "avg",
					"window": "2m"
				}
        	}
	]`
	var queries []Query
	err = json.Unmarshal([]byte(queriesData), &queries)
	assert.Nil(err)
	assert.Equal(2, len(queries))
	assert.Equal("", queries[0].Aggregator.Type)
	assert.Equal("avg", queries[1].Aggregator.Type)
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
	s := IntSeries{SeriesMeta: SeriesMeta{Tags: tags2, Name: "cpu.idle"}}
	assert.Equal(Hash(&q), Hash(&s))
}
