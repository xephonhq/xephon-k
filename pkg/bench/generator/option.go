package generator

import "time"

var DefaultOption Option
var DefaultNanosecondOption Option

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

func NewDefaultOption() Option {
	return Option{startTime: time.Now(), interval: time.Millisecond, precision: time.Millisecond}
}

func init() {
	log.Info("Init is called")
	DefaultOption = Option{startTime: time.Now(), interval: time.Millisecond, precision: time.Millisecond}
	DefaultNanosecondOption = Option{startTime: time.Now(), interval: time.Nanosecond, precision: time.Nanosecond}
}
