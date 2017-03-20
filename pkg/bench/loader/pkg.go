package loader

import (
	"fmt"
	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/util"
	"time"
)

var log = util.Logger.NewEntryWithPkg("k.b.loader")

// TODO: we should allow both config by number of points and load time, but let's use load time first
// TODO: we should allow QPS of per worker
// TODO: we should allow put multiple series in one batch, since it's allowed by many API
type Config struct {
	//TotalPoints int // 21,4748,3647
	Duration  time.Duration
	BatchSize int           // how many points we have in one batch
	WorkerNum int           // how many works
	Timeout   time.Duration // default timeout
	QPS       int           // QPS of per worker, by default is 0
	TargetDB  int
}

func (c Config) String() string {
	return fmt.Sprintf(
		`
Duration: %v
Worker number: %d
Batch size: %d
Timeout: %v
TargetDB: %s
`,
		c.Duration.Seconds(),
		c.WorkerNum,
		c.BatchSize,
		c.Timeout.Seconds(),
		bench.DBString(c.TargetDB))
}

func NewConfig(targetDB int) Config {
	return Config{
		Duration:  time.Duration(5) * time.Second,
		BatchSize: 100,
		WorkerNum: 10,
		QPS:       0,
		Timeout:   time.Duration(30) * time.Second,
		TargetDB:  targetDB,
	}
}
