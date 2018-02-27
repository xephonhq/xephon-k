package logutil

import (
	goicelog "github.com/at15/go.ice/ice/util/logutil"
	"github.com/dyweb/gommon/log"
)

var Registry = log.NewApplicationLogger()

func NewPackageLogger() *log.Logger {
	l := log.NewPackageLoggerWithSkip(1)
	Registry.AddChild(l)
	return l
}

func init() {
	Registry.AddChild(goicelog.Registry)
}
