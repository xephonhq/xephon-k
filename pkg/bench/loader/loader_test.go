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
	//ld := NewHTTPLoader(NewConfig(bench.DBXephonK), &reporter.NullReporter{}) // 8928
	//ld := NewHTTPLoader(NewConfig(bench.DBInfluxDB), &reporter.BasicReporter{})
	//INFO[0010] total request 119 pkg=k.b.reporter
	//INFO[0010] fastest 231350986 pkg=k.b.reporter
	//INFO[0010] slowest 714779592 pkg=k.b.reporter
	ld := NewHTTPLoader(NewConfig(bench.DBXephonK), &reporter.BasicReporter{}) //
	//INFO[0010] total request 11742 pkg=k.b.reporter
	//INFO[0010] fastest 1074761 pkg=k.b.reporter
	//INFO[0010] slowest 36497286 pkg=k.b.reporter
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
