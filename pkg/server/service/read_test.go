package service

import (
	"encoding/json"
	asst "github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReadServiceHTTPFactory_MakeDecode(t *testing.T) {
	assert := asst.New(t)
	queryData := `{
		"start_time": 1493363958000,
		"end_time": 1494363958000,
		"quries":[
			{
				"name":"cpi",
				"tags":{"machine":"machine-01","os":"ubuntu"},
				"match_policy": "exact",
				"start_time": 1493363958000,
				"end_time": 1494363958000
			}
		]
	}`

	var req readRequest
	err := json.NewDecoder(strings.NewReader(queryData)).Decode(&req)
	assert.Nil(err)
	t.Log(req)
	// NOTE: nothing to do with decoder
	// TODO: it seems Golang can't handle it http://stackoverflow.com/questions/21268000/unmarshaling-nested-json-objects-in-golang
	err = json.Unmarshal([]byte(queryData), &req)
	assert.Nil(err)
	t.Log(req)
}
