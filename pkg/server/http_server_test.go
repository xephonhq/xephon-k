package server_test
// NOTE: use server_test because we have cyclic import due to config

import (
	"net/http/httptest"
	"testing"

	"github.com/dyweb/gommon/requests"
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg"
	"github.com/xephonhq/xephon-k/pkg/config"
	"github.com/xephonhq/xephon-k/pkg/server/http"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"
	"github.com/xephonhq/xephon-k/pkg/util"
)

func TestHTTPServerMemoryBackendE2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skip HTTP e2e test")
	}

	util.UseTraceLog()
	util.ShowSourceLine()

	defer func() {
		util.UseDefaultLog()
		util.HideSourceLine()
	}()

	c := config.NewDaemon()
	memory.CreateStore(c.Storage.Memory)
	store, _ := memory.GetStore()
	writeService := service.NewWriteService(store)
	readService := service.NewReadService(store)
	srv := http.NewServer(c.Server.Http, writeService, readService)

	//srv := HTTPServer{Backend: "memory"}
	ts := httptest.NewServer(srv.Mux())
	defer ts.Close()

	t.Run("info", func(t *testing.T) {
		assert := asst.New(t)
		info, err := requests.GetJSONStringMap(ts.URL + "/info")
		assert.Nil(err)
		assert.Equal(pkg.Version, info["version"])
	})

	// TODO: type and precision should be string in client side instead of integers
	writeData := `[
	{
		"meta": {
			"name":"cpi",
			"type": 1,
			"precision": 1000000,
			"tags":{"machine":"machine-01","os":"ubuntu"}
		},
		"points":[[1493363958000,0],[1493363959000,1],[1493363960000,2],[1493363961000,3],[1493363962000,4]]
	},
	{
		"meta": {
			"name":"cpi",
			"type": 1,
			"precision": 1000000,
			"tags":{"machine":"machine-02","os":"ubuntu"}
		},
		"points":[[1493363958000,0],[1493363959000,1],[1493363960000,2],[1493363961000,3],[1493363962000,4]]
	},
	{
		"meta": {
			"name":"cpi",
			"type": 2,
			"precision": 1000000,
			"tags":{"machine":"machine-02","os":"ubuntu", "extra":"double"}
		},
		"points":[[1493363958000,0.2],[1493363959000,1.3],[1493363960000,2.0],[1493363961000,3.0],[1493363962000,4.1]]
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
		// FIXME: this break
		//assert.Equal(false, writeResult["error"].(bool))

		// TODO: test invalid format won't break service
		// TODO: test partially invalid payload would fail partially, I don' think it's implemented
	})

	// TODO: read, match by tag
	queryExactData := `{
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

	t.Run("exact query", func(t *testing.T) {
		assert := asst.New(t)
		res, err := requests.PostJSONString(ts.URL+"/read", queryExactData)
		assert.Nil(err)
		assert.Equal(200, res.Res.StatusCode)
		t.Log(string(res.Text))
		// TODO: validate the result instead of simply printing it
	})

	// FIXME: there is a hack in `memory/store.go` when match by filter, where name is automatically converted to __name__
	queryFilterData := `{
		"start_time": 1493363958000,
		"end_time": 1494363958000,
		"queries":[
			{
				"name":"cpi",
				"match_policy": "filter",
				"start_time": 1493363958000,
				"end_time": 1494363958000,
				"filter": {
					"type": "tag_match",
					"key": "machine",
					"value": "machine-02"
				}
			}
		]
	}`

	t.Run("filter query", func(t *testing.T) {
		assert := asst.New(t)
		res, err := requests.PostJSONString(ts.URL+"/read", queryFilterData)
		assert.Nil(err)
		assert.Equal(200, res.Res.StatusCode)
		t.Log(string(res.Text))
		// TODO: validate the result instead of simply printing it
	})

}
