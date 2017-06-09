package loader

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/client"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.bench2.worker")

type Worker struct {
	ctx    context.Context
	input  <-chan *common.IntSeries
	report chan<- *client.Result
	client *client.TSDBClient
}

func (worker *Worker) work() {
	log.Info("worker started")

	for {
		select {
		case <-worker.ctx.Done():
			log.Info("worker finished by context")
			return
		case _, ok := <-worker.input:
			if !ok {
				log.Info("worker finished by input")
				return
			}
		}
	}
}
