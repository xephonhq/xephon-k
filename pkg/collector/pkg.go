package collector

import (
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.collector")

type Collector interface {
	Update() error
}
