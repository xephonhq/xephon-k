package client

import (
	"github.com/pkg/errors"
)

type Config struct {
	Host string                 `yaml:"host"`
	Port int                    `yaml:"port"`
	URL  string                 `yaml:"url"`
	XXX  map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type configAlias Config

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
	return nil
}

func (c *Config) Validate() error {
	return nil
}
