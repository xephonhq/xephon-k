package loader

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/bench/reporter"
	"testing"
	"time"
)

func TestHTTPLoader_Run(t *testing.T) {
	if testing.Short() {
		t.Skip("skip loader run test")
	}
	//ld := NewHTTPLoader(NewConfig(bench.DBInfluxDB), &reporter.NullReporter{}) // 143
	ld := NewHTTPLoader(NewConfig(bench.DBXephonK), &reporter.NullReporter{}) // 8928
	ld.Run()
}

func testWorkWithContext(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Info("I am done!")
	}
}

func TestGoSemanticsContext(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	//defer cancel()
	testWorkWithContext(ctx)
}
