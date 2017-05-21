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

// UseDefaultLog set logger level to info
func UseDefaultLog() {
	Logger.Level = dlog.InfoLevel
	log.Info("use info logging")
}

// UseVerboseLog set logger level to debug
func UseVerboseLog() {
	Logger.Level = dlog.DebugLevel
	log.Debug("use debug logging")
}

// UseTraceLog set logger level to trace
func UseTraceLog() {
	Logger.Level = dlog.TraceLevel
	log.Trace("use trace logging")
}

func ShowSourceLine() {
	Logger.EnableSourceLine()
}

func HideSourceLine() {
	Logger.DisableSourceLine()
}
