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
	// FIXME: this is is not nanosecond, it is a problem rooted from common.IntSeries
	timeNano int64
	interval int64
}

// NewConstantValueFixedInterval returns a generator with time set to now, and interval set to 1000 nano second (1ms)
// TODO: allow configure start time
func NewConstantValueFixedInterval() ConstantValueFixedInterval {
	return ConstantValueFixedInterval{
		timeNano:  time.Now().Unix() * 1000,
		intVal:    10,
		doubleVal: 2.33,
		interval:  defaultTimeInterval,
	}
}

// GeneratorName implements Generator interface
// FIXME: this is actually string instead of name
func (gen *ConstantValueFixedInterval) GeneratorName() string {
	return fmt.Sprintf("constant generator of %d and %v", gen.intVal, gen.doubleVal)
}

// NextIntPoint implements IntGenerator interface
func (gen *ConstantValueFixedInterval) NextIntPoint() common.IntPoint {
	// NOTE: this requires using pointer receiver
	gen.timeNano += gen.interval
	return common.IntPoint{TimeNano: gen.timeNano, V: gen.intVal}
}

// NextDoublePoint implements DoubleGenerator interface
func (gen *ConstantValueFixedInterval) NextDoublePoint() common.DoublePoint {
	return common.DoublePoint{TimeNano: gen.timeNano, V: gen.doubleVal}
}
