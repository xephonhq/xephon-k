package reporter

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/client"
	"time"
)

// DiscardReporter is used to drain from channel only, it reports to nowhere
type DiscardReporter struct {
	counter int64
	start   time.Time
	end     time.Time
}

// Run implements Reporter
func (n *DiscardReporter) Run(ctx context.Context, c chan *client.Result) {
	n.start = time.Now()
LOOP:
	for {
		select {
		case <-ctx.Done():
			log.Info("null report finished by context")
			break LOOP
		case _, ok := <-c:
			// FIXED: this is never triggered?
			// The parent goroutine should sleep for a while so reporter can drain the channel
			if !ok {
				log.Info("null report finished by channel")
				break LOOP
			}
			// just drain from the channel and do nothing
			// NOTE: since null reporter is accessed by only one goroutine, this operation should be safe
			n.counter++
		}
	}
	n.end = time.Now()
}

func (n *DiscardReporter) Finalize() {
	log.Infof("run time %d s", n.end.Sub(n.start).Seconds())
	log.Infof("total request %d", n.counter)
	log.Info("null reporter has nothing to say")
}
