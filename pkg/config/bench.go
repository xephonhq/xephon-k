package config

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/util"
)

// check interface
var _ Config = (*BenchConfig)(nil)

type BenchConfig struct {
	Log  util.LogConfig         `yaml:"log" json:"log"`
	Mode string                 `yaml:"mode" json:"mode"`
	XXX  map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type benchConfigAlias BenchConfig

func NewBench() BenchConfig {
	return BenchConfig{
		Log: util.NewLogConfig(),
	}
}

func (c *BenchConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	a := (*benchConfigAlias)(c)
	if err := unmarshal(a); err != nil {
		return err
	}
	if len(c.XXX) != 0 {
		return errors.Errorf("undefined fields %v", c.XXX)
	}
	return nil
}

func (c *BenchConfig) Apply() error {
	if err := c.Validate(); err != nil {
		return err
	}
	if err := c.Log.Apply(); err != nil {
		return err
	}
	return nil
}

func (c *BenchConfig) Validate() error {
	if err := c.Log.Validate(); err != nil {
		return err
	}
	if c.Mode != "local" {
		return errors.Errorf("only support local mode, got %s instead", c.Mode)
	}
	return nil
}
