package reporter

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/bench"
)

type BasicReporter struct {
	counter int64
	fastest int64
	slowest int64
}

func (b *BasicReporter) Start(ctx context.Context, c chan *bench.RequestMetric) {
	b.slowest = 0
	b.fastest = 99999999999
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
		}
	}
}

func (b *BasicReporter) Finalize() {
	log.Infof("total request %d", b.counter)
	log.Infof("fastest %d", b.fastest)
	log.Infof("slowest %d", b.slowest)
}
