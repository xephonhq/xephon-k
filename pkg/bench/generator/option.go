package generator

import "time"

var defaultOption DefaultOption
var defaultNanosecondOption DefaultNanosecondOption

type DefaultOption struct {
}

func (opt DefaultOption) GetStartTime() time.Time {
	return time.Now()
}

func (opt DefaultOption) GetInterval() time.Duration {
	return time.Millisecond
}

func (opt DefaultOption) GetPrecision() time.Duration {
	return time.Millisecond
}

type DefaultNanosecondOption struct {
}

func (opt DefaultNanosecondOption) GetStartTime() time.Time {
	return time.Now()
}

func (opt DefaultNanosecondOption) GetInterval() time.Duration {
	return time.Nanosecond
}

func (opt DefaultNanosecondOption) GetPrecision() time.Duration {
	return time.Nanosecond
}
