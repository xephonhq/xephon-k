// Package config contains application config for daemon and bench
package config

import (
	"github.com/xephonhq/xephon-k/pkg/util"
	"gopkg.in/yaml.v2"
)

type Config interface {
	yaml.Unmarshaler
	Apply() error
	IsValid() bool
	Original() string
}

// check interface
var _ Config = (*util.LogConfig)(nil)

type DaemonConfig struct {
	Log util.LogConfig
}

func NewDaemon() *DaemonConfig {
	c := &DaemonConfig{}
	c.Log = util.NewLogConfig()
	return c
}
