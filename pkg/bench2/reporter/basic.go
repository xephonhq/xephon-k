package reporter

import (
	"context"
	"fmt"
	"github.com/VividCortex/gohistogram"
	"github.com/xephonhq/xephon-k/pkg/client"
	"time"
)

type BasicReporter struct {
	counter           int64
	start             time.Time
	end               time.Time
	fastest           int64
	slowest           int64
	totalRequestSize  int64
	totalResponseSize int64
	statusCode        map[int]int64
	hisogram          gohistogram.Histogram
}

func (b *BasicReporter) Run(ctx context.Context, c chan *client.Result) {
	b.start = time.Now()
	b.slowest = 0
	b.fastest = 99999999999
	b.statusCode = make(map[int]int64, 10)
	b.hisogram = gohistogram.NewHistogram(20)
LOOP:
	for {
		select {
		case <-ctx.Done():
			log.Info("basic report finished by context")
			break LOOP
		case result, ok := <-c:
			// FIXED: this is never triggered?
			// The parent goroutine should sleep for a while so reporter can drain the channel
			if !ok {
				log.Info("basic report finished by channel")
				break LOOP
			}
			// NOTE: since reporter is accessed by only one goroutine, these operation should be safe
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
			b.hisogram.Add(result.End.Sub(result.Start).Seconds())
		}
	}
	b.end = time.Now()
}

func (b *BasicReporter) Finalize() {
	fmt.Print(b.hisogram.String())
	log.Infof("run time %f s", b.end.Sub(b.start).Seconds())
	log.Infof("total request %d", b.counter)
	log.Infof("fastest %d", b.fastest)
	log.Infof("slowest %d", b.slowest)
	// TODO: human readable format
	log.Infof("total request size %d", b.totalRequestSize)
	log.Infof("toatl response size %d", b.totalResponseSize)
	for code, count := range b.statusCode {
		log.Infof("%d: %d", code, count)
	}
	log.Info("null reporter has nothing to say")
}
