package generator

import (
	"fmt"
	"time"

	"github.com/xephonhq/xephon-k/pkg/common"
)

var constantGeneratorDefaultTimeInterval = int64(1000)

// ConstantGenerator returns a fixed value
type ConstantGenerator struct {
	intVal    int
	doubleVal float64
	timeNano  int64
	interval  int64
}

// NewConstantGenerator returns a generator with time set to now, and interval set to 1000
func NewConstantGenerator() ConstantGenerator {
	return ConstantGenerator{
		timeNano:  time.Now().Unix() * 1000,
		intVal:    10,
		doubleVal: 2.33,
		interval:  constantGeneratorDefaultTimeInterval,
	}
}

// GeneratorName implements Generator interface
func (gen *ConstantGenerator) GeneratorName() string {
	return fmt.Sprintf("constant generator of %d and %v", gen.intVal, gen.doubleVal)
}

// NextIntPoint implements IntGenerator interface
func (gen *ConstantGenerator) NextIntPoint() common.IntPoint {
	// NOTE: this requires using pointer receiver
	gen.timeNano += gen.interval
	return common.IntPoint{TimeNano: gen.timeNano, V: gen.intVal}
}

// NextDoublePoint implements DoubleGenerator interface
func (gen *ConstantGenerator) NextDoublePoint() common.DoublePoint {
	return common.DoublePoint{TimeNano: gen.timeNano, V: gen.doubleVal}
}
