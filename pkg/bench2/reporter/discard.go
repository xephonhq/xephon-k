package reporter

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/client"
)

// DiscardReporter is used to drain from channel only, it reports to nowhere
type DiscardReporter struct {
	counter int64
}

// Start implements Reporter
func (n *DiscardReporter) Start(ctx context.Context, c chan *client.Result) {
	for {
		select {
		case <-ctx.Done():
			log.Info("null report finished via context")
			return
		case <-c:
			// just drain from the channel and do nothing
			// NOTE: since null reporter is accessed by only one goroutine, this operation should be safe
			n.counter++
		}
	}
}

func (n *DiscardReporter) Finalize() {
	log.Infof("total request %d", n.counter)
	log.Info("null reporter has nothing to say")
}
