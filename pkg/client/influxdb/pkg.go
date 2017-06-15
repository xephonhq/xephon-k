package influxdb

import (
	"net/http"

	"github.com/xephonhq/xephon-k/pkg/client"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.client.influxdb")

func New(config client.Config, transport *http.Transport) (client.TSDBClient, error) {
	return client.New(config, transport, NewSerializer())
}

func MustNew(config client.Config, transport *http.Transport) client.TSDBClient {
	c, err := New(config, transport)
	if err != nil {
		log.Panic(err)
	}
	return c
}
