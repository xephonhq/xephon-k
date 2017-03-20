package loader

import (
	"context"
	"net/http"
	"sync"

	"fmt"
	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/bench/reporter"
	"github.com/xephonhq/xephon-k/pkg/server"
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
	// NOTE: https://github.com/at15/mini-impl/issues/1
	tr := &http.Transport{
		MaxIdleConnsPerHost: loader.config.WorkerNum,
	}
	log.Infof("target db %s", bench.DBString(loader.config.TargetDB))
	switch loader.config.TargetDB {
	case bench.DBInfluxDB:
		req, err := http.NewRequest("POST", "http://localhost:8086/write?db=sb", nil)
		if err != nil {
			log.Panic(err)
			return
		}
		baseReq = req
	case bench.DBXephonK:
		url := fmt.Sprintf("http://localhost:%d/write", server.DefaultPort)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			log.Panic(err)
			return
		}
		baseReq = req
	case bench.DBKairosDB:
		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/datapoints", nil)
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
	// NOTE: for basic reporter, data is not written to TSDB, no need to wait actually ...
	reportCtx, cancelReport := context.WithTimeout(context.Background(),
		loader.config.Duration+time.Duration(1)*time.Second)
	defer cancelReport() // FIXME: this should be useless
	go func() {
		loader.report.Start(reportCtx, loader.metricChan)
		wg.Done()
	}()
	wg.Wait()
	loader.report.Finalize()
	log.Info("loader finished")
}
