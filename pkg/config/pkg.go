// Package config contains application config for daemon and bench
package config

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/server"
	"github.com/xephonhq/xephon-k/pkg/server/grpc"
	"github.com/xephonhq/xephon-k/pkg/server/http"
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
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
var _ Config = (*DaemonConfig)(nil)
var _ Config = (*util.LogConfig)(nil)
var _ Config = (*storage.Config)(nil)
var _ Config = (*memory.Config)(nil)
var _ Config = (*disk.Config)(nil)
var _ Config = (*cassandra.Config)(nil)
var _ Config = (*server.Config)(nil)
var _ Config = (*http.Config)(nil)
var _ Config = (*grpc.Config)(nil)

type DaemonConfig struct {
	Log     util.LogConfig         `yaml:"log" json:"log"`
	Storage storage.Config         `yaml:"storage" json:"storage"`
	Server  server.Config          `yaml:"server" json:"server"`
	XXX     map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type daemonConfigAlias DaemonConfig

func NewDaemon() DaemonConfig {
	return DaemonConfig{
		Log:     util.NewLogConfig(),
		Storage: storage.NewConfig(),
		Server:  server.NewConfig(),
	}
}

func (c *DaemonConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	a := (*daemonConfigAlias)(c)
	if err := unmarshal(a); err != nil {
		return err
	}
	if len(c.XXX) != 0 {
		return errors.Errorf("undefined fields %v", c.XXX)
	}
	return nil
}

func (c *DaemonConfig) Apply() error {
	if err := c.Validate(); err != nil {
		return err
	}
	if err := c.Log.Apply(); err != nil {
		return err
	}
	if err := c.Storage.Apply(); err != nil {
		return err
	}
	if err := c.Server.Apply(); err != nil {
		return err
	}
	return nil
}

func (c *DaemonConfig) Validate() error {
	if err := c.Log.Validate(); err != nil {
		return err
	}
	if err := c.Storage.Validate(); err != nil {
		return err
	}
	if err := c.Server.Validate(); err != nil {
		return err
	}
	return nil
}
