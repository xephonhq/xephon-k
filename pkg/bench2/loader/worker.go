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
	client client.TSDBClient
}

func NewWorker(ctx context.Context, input <-chan *common.IntSeries, report chan<- *client.Result, client client.TSDBClient) *Worker {
	return &Worker{
		ctx:    ctx,
		input:  input,
		report: report,
		client: client,
	}
}

func (worker *Worker) Work() {
	log.Info("worker started")

	for {
		select {
		case <-worker.ctx.Done():
			log.Info("worker finished by context")
			return
		case s, ok := <-worker.input:
			// TODO: make the real request
			worker.client.WriteInt(s)
			if !ok {
				log.Info("worker finished by input")
				return
			}
		}
	}
}
