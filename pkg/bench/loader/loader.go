package loader

import "time"

// TODO: we should allow both config by number of points and load time, but let's use load time first
// TODO: we should allow put multiple series in one batch, since it's allowed by many API
type Config struct {
	//TotalPoints int // 21,4748,3647
	Duration  time.Duration
	BatchSize int           // how many points we have in one batch
	WorkerNum int           // how many works
	Timeout   time.Duration // default timeout
	QPS       int           // QPS of per worker, by default is 0
}

type Loader struct {
	config Config
}

func NewConfig() Config {
	return Config{
		Duration:  time.Duration(1) * time.Minute,
		BatchSize: 100,
		WorkerNum: 10,
		QPS:       0,
		Timeout:   time.Duration(30) * time.Second,
	}
}

func (loader *Loader) Run() {

}
