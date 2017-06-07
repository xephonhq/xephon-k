package http

import (
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.server.http")

type Server struct {
	store storage.Store
}
