package generator

import (
	"github.com/pkg/errors"
)

type Config struct {
	TimeInterval    int                    `yaml:"timeInterval"`
	TimeNoise       bool                   `yaml:"timeNoise"`
	PointsPerSeries int                    `yaml:"pointsPerSeries"`
	NumSeries       int                    `yaml:"numSeries"`
	XXX             map[string]interface{} `yaml:",inline"`
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
