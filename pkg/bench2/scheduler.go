package bench2

import (
	"github.com/xephonhq/xephon-k/pkg/client"
	"github.com/xephonhq/xephon-k/pkg/client/xephonk"
	"github.com/xephonhq/xephon-k/pkg/config"
	"net/http"
	"sync"
)

type Scheduler struct {
	config config.BenchConfig
}

func NewScheduler(config config.BenchConfig) *Scheduler {
	return &Scheduler{
		config: config,
	}
}

func (scheduler *Scheduler) Run() error {
	transport := &http.Transport{
		MaxIdleConns:        scheduler.config.Loader.WorkerNum,
		MaxIdleConnsPerHost: scheduler.config.Loader.WorkerNum,
	}

	var (
		c client.TSDBClient
		//err error
	)
	switch scheduler.config.Loader.Target {
	case "xephonk":
		c = xephonk.MustNew(scheduler.config.Targets.XephonK, transport)
	default:
		log.Fatal("only support xephonk for now")
	}
	// TODO: remove me, for the sake of no compile error
	c.Send()

	var wg sync.WaitGroup
	wg.Add(scheduler.config.Loader.WorkerNum)
	for i := 0; i < scheduler.config.Loader.WorkerNum; i++ {
		go func() {
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
