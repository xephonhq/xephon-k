package xephonk

import (
	"net/http"

	"github.com/xephonhq/xephon-k/pkg/client"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.client.xephonk")

func New(config client.Config, transport *http.Transport) (client.TSDBClient, error) {
	return client.New(config, transport, NewSeializer())
}
