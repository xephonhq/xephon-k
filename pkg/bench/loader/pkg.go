package loader

import (
	"github.com/xephonhq/xephon-k/pkg/util"
	"time"
)

var log = util.Logger.NewEntryWithPkg("k.b.loader")

// TODO: we should allow both config by number of points and load time, but let's use load time first
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

func NewConfig(targetDB int) Config {
	return Config{
		//Duration:  time.Duration(1) * time.Minute,
		Duration:  time.Duration(5) * time.Second,
		BatchSize: 100,
		WorkerNum: 10,
		QPS:       0,
		Timeout:   time.Duration(30) * time.Second,
		TargetDB:  targetDB,
	}
}

type result struct {
	err      error
	code     int
	duration time.Duration
}
