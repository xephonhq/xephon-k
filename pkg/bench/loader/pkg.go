package loader

import (
	"github.com/xephonhq/xephon-k/pkg/util"
	"time"
)

var log = util.Logger.NewEntryWithPkg("k.b.loader")

type result struct {
	err      error
	code     int
	duration time.Duration
}
