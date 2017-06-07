// Package config contains application config for daemon and bench
package config

import (
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/storage/disk"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"
	"github.com/xephonhq/xephon-k/pkg/util"
	"gopkg.in/yaml.v2"
)

type Config interface {
	yaml.Unmarshaler
	Apply() error
	Validate() error
	// TODO: we can't have original because YAML does not have []byte like in JSON for Unmarshaler
	// TODO: we call validate in apply, which trigger validate of all the children, then the apply of all the children
	// also trigger the validate of their own and their children, the lowest level would have its validate called n times
	// where n is their nested level
}

// check interface
var _ Config = (*util.LogConfig)(nil)
var _ Config = (*storage.Config)(nil)
var _ Config = (*memory.Config)(nil)
var _ Config = (*disk.Config)(nil)

type DaemonConfig struct {
	Log     util.LogConfig `yaml:"log" json:"log"`
	Storage storage.Config `yaml:"storage" json:"storage"`
}

func NewDaemon() *DaemonConfig {
	c := &DaemonConfig{
		Log:     util.NewLogConfig(),
		Storage: storage.NewConfig(),
	}
	return c
}
