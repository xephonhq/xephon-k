package generator

import (
	"fmt"
	"time"

	"github.com/xephonhq/xephon-k/pkg/common"
)

// ConstantGenerator returns a fixed value
type ConstantGenerator struct {
	intVal    int
	doubleVal float64
}

// GeneratorName implements Generator interface
func (gen ConstantGenerator) GeneratorName() string {
	return fmt.Sprintf("constant generator of %d and %v", gen.intVal, gen.doubleVal)
}

// NextIntPoint implements IntGenerator interface
// TODO: use pointer receiver?
func (gen ConstantGenerator) NextIntPoint() common.IntPoint {
	// TODO: there will be duplicate time if we use real time, should use atomic?
	return common.IntPoint{TimeNano: time.Now().Unix() * 1000, V: gen.intVal}
}

// NextDoublePoint implements DoubleGenerator interface
func (gen ConstantGenerator) NextDoublePoint() common.DoublePoint {
	return common.DoublePoint{TimeNano: time.Now().Unix() * 1000, V: gen.doubleVal}
}
