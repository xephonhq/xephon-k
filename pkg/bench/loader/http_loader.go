package loader

import (
	"context"
	"net/http"
	"sync"

	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/bench/reporter"
	"time"
)

type HTTPLoader struct {
	config     Config
	metricChan chan *bench.RequestMetric
	report     reporter.Reporter
}

func NewHTTPLoader(config Config, report reporter.Reporter) *HTTPLoader {
	return &HTTPLoader{
		config:     config,
		metricChan: make(chan *bench.RequestMetric, config.WorkerNum*10), // TODO: proper size of channel
		report:     report,
	}
}

func (loader *HTTPLoader) Run() {
	var baseReq *http.Request
	tr := &http.Transport{}
	log.Info(loader.config.TargetDB)
	switch loader.config.TargetDB {
	case bench.DBInfluxDB:
		req, err := http.NewRequest("POST", "http://localhost:8086/write?db=sb", nil)
		if err != nil {
			log.Panic(err)
			return
		}
		baseReq = req
	case bench.DBXephonK:
		req, err := http.NewRequest("POST", "http://localhost:8080/write", nil)
		if err != nil {
			log.Panic(err)
			return
		}
		baseReq = req
	default:
		log.Panic("unsupported database, no base request avaliable")
		return
	}

	var wg sync.WaitGroup
	wg.Add(loader.config.WorkerNum)
	for i := 0; i < loader.config.WorkerNum; i++ {
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), loader.config.Duration)
			// TODO: I don't think this cancel will be called
			defer cancel()
			worker := NewHTTPWorker(loader.config, ctx, baseReq, tr, loader.metricChan)
			worker.work()
			wg.Done()
		}()
	}
	wg.Add(1)
	// TODO: better control of report ending
	reportCtx, cancelReport := context.WithTimeout(context.Background(),
		loader.config.Duration+time.Duration(5)*time.Second)
	defer cancelReport() // FIXME: this should be useless
	go func() {
		loader.report.Start(reportCtx, loader.metricChan)
		wg.Done()
	}()
	wg.Wait()
	loader.report.Finalize()
	log.Info("loader finished")
}
