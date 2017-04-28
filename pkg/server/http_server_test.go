package server

import (
	"encoding/json"
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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
		// TODO: add util function, following block is pretty verbose
		res, err := http.Get(ts.URL + "/info")
		assert.Nil(err)
		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		assert.Nil(err)
		var info map[string]string
		err = json.Unmarshal(data, &info)
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
		"tags":{"machine":"machine-01","os":"ubuntu"},
		"points":[[1493363958000,0],[1493363959000,1],[1493363960000,2],[1493363961000,3],[1493363962000,4]]
	}
	]`

	t.Run("write", func(t *testing.T) {
		assert := asst.New(t)
		// TODO: add util function, following block is pretty verbose
		res, err := http.Post(ts.URL+"/write", "application/json", strings.NewReader(writeData))
		assert.Nil(err)
		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		assert.Nil(err)
		var writeResult map[string]interface{}
		err = json.Unmarshal(data, &writeResult)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode)
		assert.Equal(false, writeResult["error"].(bool))

		// TODO: test invalid format won't break service
	})

	// TODO: read

}
