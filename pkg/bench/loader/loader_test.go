package loader

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/bench"
	"testing"
	"time"
)

func TestHTTPLoader_Run(t *testing.T) {
	if testing.Short() {
		t.Skip("skip loader run test")
	}
	//ld := NewHTTPLoader(NewConfig(bench.DBInfluxDB))
	ld := NewHTTPLoader(NewConfig(bench.DBXephonK))
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
