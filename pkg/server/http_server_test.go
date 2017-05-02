package server

import (
	"github.com/dyweb/gommon/requests"
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg"
	"net/http/httptest"
	"testing"
)

func TestHTTPServerMemoryBackendE2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skip HTTP e2e test")
	}

	srv := HTTPServer{Backend: "memory"}
	ts := httptest.NewServer(srv.Mux())
	defer ts.Close()

	t.Run("info", func(t *testing.T) {
		assert := asst.New(t)
		info, err := requests.GetJSONStringMap(ts.URL + "/info")
		assert.Nil(err)
		assert.Equal(pkg.Version, info["version"])
	})

	writeData := `[
	{
		"name":"cpi",
		"tags":{"machine":"machine-01","os":"ubuntu"},
		"points":[[1493363958000,0],[1493363959000,1],[1493363960000,2],[1493363961000,3],[1493363962000,4]]
	},
	{
		"name":"cpi",
		"tags":{"machine":"machine-02","os":"ubuntu"},
		"points":[[1493363958000,0],[1493363959000,1],[1493363960000,2],[1493363961000,3],[1493363962000,4]]
	}
	]`

	t.Run("write", func(t *testing.T) {
		assert := asst.New(t)
		res, err := requests.PostJSONString(ts.URL+"/write", writeData)
		assert.Nil(err)
		var writeResult map[string]interface{}
		err = res.JSON(&writeResult)
		assert.Nil(err)
		assert.Equal(200, res.Res.StatusCode)
		assert.Equal(false, writeResult["error"].(bool))

		// TODO: test invalid format won't break service
		// TODO: test partially invalid payload would fail partially, I don' think it's implemented
	})

	// TODO: read
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

	t.Run("exact query", func(t *testing.T) {
		assert := asst.New(t)
		res, err := requests.PostJSONString(ts.URL+"/read", queryData)
		assert.Nil(err)
		// FIXME: it's 500
		t.Log(string(res.Text))
		// FIXME: there is no logging
		assert.Equal(200, res.Res.StatusCode)
	})

}
