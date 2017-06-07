package util

import (
	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"
)

// Log util and config

// Logger is the default logger with info level
var Logger = dlog.NewLogger()

// Short name use in util package
var log = Logger.NewEntryWithPkg("k.util")

type LogConfig struct {
	Level  string                 `yaml:"level" json:"level"`
	Color  bool                   `yaml:"color" json:"color"`
	Source bool                   `yaml:"source" json:"source"`
	XXX    map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type logConfigAlias LogConfig

func NewLogConfig() LogConfig {
	return LogConfig{
		Level:  "info",
		Color:  true,
		Source: false,
	}
}

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

func (c *LogConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	a := (*logConfigAlias)(c)
	if err := unmarshal(a); err != nil {
		return err
	}
	if len(c.XXX) != 0 {
		return errors.Errorf("undefined fields %v", c.XXX)
	}
	return nil
}

func (c *LogConfig) Apply() error {
	if err := c.Validate(); err != nil {
		return err
	}
	if Logger.Level.String() != c.Level {
		newLevel, err := dlog.ParseLevel(c.Level, false)
		if err != nil {
			return errors.Wrapf(err, "can't set logging level to %s", c.Level)
		}
		Logger.Level = newLevel
	}
	// TODO: handle color, formatter interface does not expose this functionality
	if c.Source {
		ShowSourceLine()
	} else {
		HideSourceLine()
	}
	return nil
}

func (c *LogConfig) Validate() error {
	if _, err := dlog.ParseLevel(c.Level, false); err != nil {
		return errors.Wrap(err, "invalid log config")
	}
	return nil
}
