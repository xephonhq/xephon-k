package storage

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
	"github.com/xephonhq/xephon-k/pkg/storage/disk"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"
)

type Config struct {
	Memory    memory.Config          `yaml:"memory" json:"memory"`
	Disk      disk.Config            `yaml:"disk" json:"disk"`
	Cassandra cassandra.Config       `yaml:"cassandra" json:"cassandra"`
	XXX       map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type configAlias Config

func NewConfig() Config {
	return Config{
		Memory:    memory.NewConfig(),
		Disk:      disk.NewConfig(),
		Cassandra: cassandra.NewConfig(),
	}
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	a := (*configAlias)(c)
	if err := unmarshal(a); err != nil {
		return err
	}
	if len(c.XXX) != 0 {
		return errors.Errorf("undefined fields %v", c.XXX)
	}
	return nil
}

func (c *Config) Apply() error {
	if err := c.Validate(); err != nil {
		return err
	}
	if err := c.Memory.Apply(); err != nil {
		return err
	}
	if err := c.Disk.Apply(); err != nil {
		return err
	}
	if err := c.Cassandra.Apply(); err != nil {
		return err
	}
	return nil
}

func (c *Config) Validate() error {
	if err := c.Memory.Validate(); err != nil {
		return err
	}
	if err := c.Disk.Validate(); err != nil {
		return err
	}
	if err := c.Cassandra.Validate(); err != nil {
		return err
	}
	return nil
}
