package generator

import (
	"fmt"
	"time"

	"github.com/xephonhq/xephon-k/pkg/common"
)

// ConstantValueFixedInterval returns a constant value with timestamp increasing at interval
type ConstantValueFixedInterval struct {
	intVal    int
	doubleVal float64
	timestamp int64 // timestamp is in different precision based on the option. see https://github.com/xephonhq/xephon-k/issues/35
	interval  int64 // interval is in same precision with timestamp
	option    Option
}

// NewConstantValueFixedInterval returns a generator with time set to now, and interval set to 1000 nano second (1ms)
func NewConstantValueFixedInterval(option Option) ConstantValueFixedInterval {
	// NOTE: we use int64 instead of time.Time to avoid the overhead of object creation from Add(), which creates a new time.Time
	var startTime int64
	var interval int64

	switch option.GetPrecision() {
	case time.Second:
		startTime = option.GetStartTime().Unix()
		interval = option.GetInterval().Nanoseconds() / 1000000000
	case time.Millisecond:
		startTime = option.GetStartTime().Unix() * 1000
		interval = option.GetInterval().Nanoseconds() / 1000000
	case time.Nanosecond:
		startTime = option.GetStartTime().UnixNano()
		interval = option.GetInterval().Nanoseconds()
	default:
		log.Panicf("unsupported precision %v", option.GetPrecision())
		return ConstantValueFixedInterval{}
	}
	return ConstantValueFixedInterval{
		timestamp: startTime,
		interval:  interval,
		intVal:    10,
		doubleVal: 2.33,
		option:    option,
	}
}

func (gen *ConstantValueFixedInterval) String() string {
	return fmt.Sprintf("generator: constant value of %d and %v, fixed interval %v start from %v in precision of %v",
		gen.intVal, gen.doubleVal, gen.option.GetInterval(), gen.option.GetStartTime(), gen.option.GetPrecision())
}

// GetOption implements Generator interface
func (gen *ConstantValueFixedInterval) GetOption() Option {
	return gen.option
}

// NextIntPoint implements IntGenerator interface
func (gen *ConstantValueFixedInterval) NextIntPoint() common.IntPoint {
	// NOTE: this requires using pointer receiver
	gen.timestamp += gen.interval
	return common.IntPoint{TimeNano: gen.timestamp, V: gen.intVal}
}

// NextDoublePoint implements DoubleGenerator interface
func (gen *ConstantValueFixedInterval) NextDoublePoint() common.DoublePoint {
	gen.timestamp += gen.interval
	return common.DoublePoint{TimeNano: gen.timestamp, V: gen.doubleVal}
}
