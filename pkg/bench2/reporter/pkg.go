package reporter

import (
	"context"

	"github.com/xephonhq/xephon-k/pkg/client"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.bench2.reporter")

// check interface

var _ Reporter = (*DiscardReporter)(nil)

type Reporter interface {
	//ReporterName() string
	Start(context.Context, chan *client.Result)
	Finalize()
}
