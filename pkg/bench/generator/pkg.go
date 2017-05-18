package generator

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("xkb.generator")

// check interface
var _ IntGenerator = (*ConstantValueFixedInterval)(nil)
var _ DoubleGenerator = (*ConstantValueFixedInterval)(nil)

type Generator interface {
	GetOption() Option
}

// TODO: maybe using a concrete type is better
// type Option interface {
// 	GetStartTime() time.Time
// 	GetInterval() time.Duration
// 	GetPrecision() time.Duration
// }

type IntGenerator interface {
	Generator
	NextIntPoint() common.IntPoint // TODO: use pointer instead of value?
}

type DoubleGenerator interface {
	Generator
	NextDoublePoint() common.DoublePoint
}
