package reporter

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/client"
)

// NullReporter is used to drain from channel only, it reports to nowhere
type NullReporter struct {
	counter int64
}

// Start implements Reporter
func (n *NullReporter) Start(ctx context.Context, c chan *client.Result) {
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

func (n *NullReporter) Finalize() {
	log.Infof("total request %d", n.counter)
	log.Info("null reporter has nothing to say")
}
