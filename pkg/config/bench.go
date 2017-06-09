package config

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/bench2/generator"
	"github.com/xephonhq/xephon-k/pkg/bench2/loader"
	"github.com/xephonhq/xephon-k/pkg/client"
	"github.com/xephonhq/xephon-k/pkg/util"
)

// check interface
var _ Config = (*BenchConfig)(nil)

type BenchConfig struct {
	Log       util.LogConfig         `yaml:"log" json:"log"`
	Mode      string                 `yaml:"mode" json:"mode"`
	Loader    loader.Config          `yaml:"loader" json:"loader"`
	Generator generator.Config       `yaml:"generator"`
	Targets   Targets                `yaml:"targets"`
	XXX       map[string]interface{} `yaml:",inline"`
}

type Targets struct {
	InfluxDB client.Config `yaml:"influxdb"`
	XephonK  client.Config `yaml:"xephonk"`
	KairosDB client.Config `yaml:"kairosdb"`
}

// avoid recursion in UnmarshalYAML
type benchConfigAlias BenchConfig

func NewBench() BenchConfig {
	return BenchConfig{
		Log:    util.NewLogConfig(),
		Mode:   "local",
		Loader: loader.NewConfig(),
		Targets: Targets{
			InfluxDB: client.Config{
				Host:    "localhost",
				Port:    8086,
				URL:     "write?db=sb",
				Timeout: 30,
			}, XephonK: client.Config{
				Host:    "localhost",
				Port:    2333,
				URL:     "write",
				Timeout: 30,
			}, KairosDB: client.Config{
				Host:    "localhost",
				Port:    8080,
				URL:     "api/v1/datapoints",
				Timeout: 30,
			}},
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
