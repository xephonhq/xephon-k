package loader

import (
	"context"
	"net/http"
	"sync"

	"github.com/xephonhq/xephon-k/pkg/bench"
)

type HTTPLoader struct {
	config Config
}

func NewHTTPLoader(config Config) *HTTPLoader {
	return &HTTPLoader{config: config}
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
			worker := NewHTTPWorker(loader.config, ctx, baseReq, tr)
			worker.work()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Info("loader finished")
}
