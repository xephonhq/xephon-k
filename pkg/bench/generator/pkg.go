package generator

import "github.com/xephonhq/xephon-k/pkg/common"

// check interface
var _ IntGenerator = (*ConstantGenerator)(nil)
var _ DoubleGenerator = (*ConstantGenerator)(nil)

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
