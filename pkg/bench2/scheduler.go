package bench2

import (
	"context"
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/bench2/generator"
	"github.com/xephonhq/xephon-k/pkg/bench2/loader"
	"github.com/xephonhq/xephon-k/pkg/bench2/reporter"
	"github.com/xephonhq/xephon-k/pkg/client"
	"github.com/xephonhq/xephon-k/pkg/client/influxdb"
	"github.com/xephonhq/xephon-k/pkg/client/kairosdb"
	"github.com/xephonhq/xephon-k/pkg/client/xephonk"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/config"
	"net/http"
	"sync"
	"time"
)

type Scheduler struct {
	config   config.BenchConfig
	reporter reporter.Reporter
}

func NewScheduler(config config.BenchConfig) (*Scheduler, error) {
	s := &Scheduler{
		config: config,
	}
	switch config.Loader.Reporter {
	case "discard", "null":
		s.reporter = &reporter.DiscardReporter{}
	case "basic":
		s.reporter = &reporter.BasicReporter{}
	default:
		return nil, errors.Errorf("unsupported reporter", config.Loader.Reporter)
	}
	return s, nil
}

func (scheduler *Scheduler) Run() error {
	var ctx context.Context

	switch scheduler.config.Loader.LimitBy {
	case "time":
		ctx, _ = context.WithTimeout(context.Background(), time.Duration(scheduler.config.Loader.Time)*time.Second)
	case "points":
		ctx = context.Background()
	}

	gen := generator.NewConstantGenerator(scheduler.config.Generator)
	data := make(chan *common.IntSeries, scheduler.config.Loader.WorkerNum)
	report := make(chan *client.Result, scheduler.config.Loader.WorkerNum)

	transport := &http.Transport{
		MaxIdleConns:        scheduler.config.Loader.WorkerNum,
		MaxIdleConnsPerHost: scheduler.config.Loader.WorkerNum,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if "points" == scheduler.config.Loader.LimitBy {
			for i := 0; i < scheduler.config.Loader.Points; {
				data <- gen.NextInt()
				i += scheduler.config.Generator.PointsPerSeries
			}
			log.Infof("generator stopped after %d points", scheduler.config.Loader.Points)
			goto GENERATOR_FINISH
		}
		if "time" == scheduler.config.Loader.LimitBy {
			for {
				select {
				case <-ctx.Done():
					log.Infof("generator stopped after %d seconds", scheduler.config.Loader.Time)
					// NOTE: break will only break the select, thus we use goto
					// https://stackoverflow.com/questions/11104085/in-go-does-a-break-statement-break-from-a-switch-select
					goto GENERATOR_FINISH
				default:
					data <- gen.NextInt()
				}
			}
		}
	GENERATOR_FINISH:
		log.Info("close data channel")
		close(data)
		wg.Done()
	}()
	wg.Add(scheduler.config.Loader.WorkerNum)
	for i := 0; i < scheduler.config.Loader.WorkerNum; i++ {
		go func() {
			var c client.TSDBClient
			switch scheduler.config.Loader.Target {
			case "xephonk":
				c = xephonk.MustNew(scheduler.config.Targets.XephonK, transport)
			case "influxdb":
				c = influxdb.MustNew(scheduler.config.Targets.InfluxDB, transport)
			case "kairosdb":
				c = kairosdb.MustNew(scheduler.config.Targets.KairosDB, transport)
			default:
				log.Fatal("only support xephonk for now")
			}
			worker := loader.NewWorker(ctx, data, report, c)
			worker.Work()
			wg.Done()
		}()
	}

	go func() {
		scheduler.reporter.Run(context.Background(), report)
	}()
	wg.Wait()
	// A closed channel can still be drained from https://dave.cheney.net/tag/golang-3
	close(report)
	// NOTE: sleep 1s so the reporter can drained the result channel
	time.Sleep(time.Second)
	scheduler.reporter.Finalize()
	return nil
}
