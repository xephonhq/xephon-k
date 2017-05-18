package generator

import "github.com/xephonhq/xephon-k/pkg/common"

// default time interval for fixed time interval generators in nanoseconds
var defaultTimeInterval = int64(1000)

// check interface
var _ IntGenerator = (*ConstantValueFixedInterval)(nil)
var _ DoubleGenerator = (*ConstantValueFixedInterval)(nil)

type Generator interface {
	GeneratorName() string
}

type IntGenerator interface {
	Generator
	NextIntPoint() common.IntPoint // TODO: use pointer instead of value?
}

type DoubleGenerator interface {
	Generator
	NextDoublePoint() common.DoublePoint
}
