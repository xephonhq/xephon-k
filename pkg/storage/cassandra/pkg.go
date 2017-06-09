package cassandra

import (
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.storage.cassandra")

var defaultKeySpace = "xephon"
