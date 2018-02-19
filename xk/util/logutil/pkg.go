package logutil

import (
	"github.com/dyweb/gommon/log"
)

var Registry = log.NewApplicationLogger()

func NewPackageLogger() *log.Logger {
	l := log.NewPackageLoggerWithSkip(1)
	Registry.AddChild(l)
	return l
}
