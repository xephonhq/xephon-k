package http

import (
	"encoding/json"
	asst "github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// TODO: rename, it's copied from old read service
func TestReadServiceHTTPFactory_MakeDecode(t *testing.T) {
	assert := asst.New(t)
	queryData := `{
		"start_time": 1493363958000,
		"end_time": 1494363958000,
		"queries":[
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
	// NOTE: nothing to do with decoder
	// FIXED: Golang can't handle nested struct array http://stackoverflow.com/questions/21268000/unmarshaling-nested-json-objects-in-golang
	err = json.Unmarshal([]byte(queryData), &req)
	assert.Nil(err)
	assert.Equal(1, len(req.Queries))
	assert.Equal("cpi", req.Queries[0].Name)
}
