package reporter

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/bench"
)

type BasicReporter struct {
	counter           int64
	fastest           int64
	slowest           int64
	totalRequestSize  int64
	totalResponseSize int64
	statusCode        map[int]int64
}

func (b *BasicReporter) Start(ctx context.Context, c chan *bench.RequestMetric) {
	b.slowest = 0
	b.fastest = 99999999999
	b.statusCode = make(map[int]int64, 10)
	for {
		select {
		case <-ctx.Done():
			log.Info("basic report finished via context")
			return
		case result := <-c:
			d := result.End.Sub(result.Start).Nanoseconds()
			if d < b.fastest {
				b.fastest = d
			}
			if d > b.slowest {
				b.slowest = d
			}
			b.counter++
			b.totalRequestSize += result.RequestSize
			b.totalResponseSize += result.ResponseSize
			// TODO: if the key does not exist, the value should be 0?
			b.statusCode[result.Code] += 1
		}
	}
}

func (b *BasicReporter) Finalize() {
	log.Infof("total request %d", b.counter)
	log.Infof("fastest %d", b.fastest)
	log.Infof("slowest %d", b.slowest)
	// TODO: human readable format
	log.Infof("total request size %d", b.totalRequestSize)
	log.Infof("toatl response size %d", b.totalResponseSize)
	for code, count := range b.statusCode {
		log.Infof("%d: %d", code, count)
	}
}
