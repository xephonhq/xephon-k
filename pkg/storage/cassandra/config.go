package cassandra

import (
	"github.com/pkg/errors"
)

// TODO: keyspace
type Config struct {
	Host     string                 `yaml:"host" json:"host"`
	Port     int                    `yaml:"port" json:"port"`
	Keyspace string                 `yaml:"keyspace" json:"keyspace"`
	XXX      map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type configAlias Config

func NewConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     9042,
		Keyspace: defaultKeySpace,
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
	return nil
}

func (c *Config) Validate() error {
	// TODO: try to connect to cassandra?
	if c.Port < 0 {
		return errors.Errorf("port number must be positive, got %d ", c.Port)
	}
	return nil
}
