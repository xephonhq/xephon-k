package generator

import "time"

var DefaultSecondOption = Option{startTime: time.Now(), interval: time.Second, precision: time.Second}

// DefaultOption is Millisecond
var DefaultOption = Option{startTime: time.Now(), interval: time.Millisecond, precision: time.Millisecond}
var DefaultNanosecondOption = Option{startTime: time.Now(), interval: time.Nanosecond, precision: time.Nanosecond}

type Option struct {
	startTime time.Time
	interval  time.Duration
	precision time.Duration
}

func (opt Option) GetStartTime() time.Time {
	return opt.startTime
}

func (opt Option) GetInterval() time.Duration {
	return opt.interval
}

func (opt Option) GetPrecision() time.Duration {
	return opt.precision
}
