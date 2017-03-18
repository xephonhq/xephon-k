package reporter

import (
	"context"

	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.b.reporter")

// check interface

var _ Reporter = (*NullReporter)(nil)
var _ Reporter = (*NullReporter)(nil)

type Reporter interface {
	//ReporterName() string
	Start(context.Context, chan *bench.RequestMetric)
	Finalize()
}
