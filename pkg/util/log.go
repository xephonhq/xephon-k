package util

import dlog "github.com/dyweb/gommon/log"

// Log util

// Logger is the default logger with info level
var Logger = dlog.NewLogger()

// Short name use in util package
var log = Logger.NewEntryWithPkg("k.util")

func init() {
	f := dlog.NewTextFormatter()
	f.EnableColor = true
	Logger.Formatter = f
	Logger.Level = dlog.InfoLevel
}

// UseVerboseLog set logger level to debug
func UseVerboseLog() {
	Logger.Level = dlog.DebugLevel
	log.Debug("enable debug logging")
}
