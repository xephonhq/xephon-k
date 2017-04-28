package loader

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/bench/reporter"
	"testing"
	"time"
)

func TestHTTPLoader_Run(t *testing.T) {
	t.Skip("skip loader run test")
	//ld := NewHTTPLoader(NewConfig(bench.DBInfluxDB), &reporter.NullReporter{}) // 143
	//ld := NewHTTPLoader(NewConfig(bench.DBXephonK), &reporter.NullReporter{}) // 8928
	//ld := NewHTTPLoader(NewConfig(bench.DBInfluxDB), &reporter.BasicReporter{})
	//INFO[0010] total request 119 pkg=k.b.reporter
	//INFO[0010] fastest 231350986 pkg=k.b.reporter
	//INFO[0010] slowest 714779592 pkg=k.b.reporter
	//ld := NewHTTPLoader(NewConfig(bench.DBXephonK), &reporter.BasicReporter{}) //
	//INFO[0010] total request 11742 pkg=k.b.reporter
	//INFO[0010] fastest 1074761 pkg=k.b.reporter
	//INFO[0010] slowest 36497286 pkg=k.b.reporter

	// try with local influxdb
	//ld := NewHTTPLoader(NewConfig(bench.DBInfluxDB), &reporter.BasicReporter{})
	// when db is not created, it seems pretty good
	//INFO[0010] total request 14446 pkg=k.b.reporter
	//INFO[0010] fastest 231121 pkg=k.b.reporter
	//INFO[0010] slowest 48644323 pkg=k.b.reporter
	// well for local
	//INFO[0010] total request 125 pkg=k.b.reporter
	//INFO[0010] fastest 192001022 pkg=k.b.reporter
	//INFO[0010] slowest 549661074 pkg=k.b.reporter
	//INFO[0010] loader finished pkg=k.b.loader

	ld := NewHTTPLoader(NewConfig(bench.DBKairosDB), &reporter.BasicReporter{})
	//INFO[0010] basic report finished via context pkg=k.b.reporter
	//INFO[0010] total request 14191 pkg=k.b.reporter
	//INFO[0010] fastest 237171 pkg=k.b.reporter
	//INFO[0010] slowest 349506588 pkg=k.b.reporter
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
