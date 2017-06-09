package loader

import (
	"github.com/pkg/errors"
)

type Config struct {
	Target        string                 `yaml:"target" json:"target"`
	LimitBy       string                 `yaml:"limitBy" json:"limit_by"`
	Points        int                    `yaml:"points" json:"points"`
	Series        int                    `yaml:"series" json:"series"`
	Time          int                    `yaml:"time" json:"time"`
	WorkerNum     int                    `yaml:"workerNum" json:"workerNum"`
	WorkerTimeout int                    `yaml:"workerTimeout" json:"worker_timeout"`
	XXX           map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type configAlias Config

func NewConfig() Config {
	return Config{
		Target:        "xephonk",
		LimitBy:       "time",
		Points:        1000000,
		Series:        1,
		Time:          10,
		WorkerNum:     10,
		WorkerTimeout: 30,
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
	// TODO: real apply
	return nil
}

func (c *Config) Validate() error {
	// TODO: real validate
	return nil
}
