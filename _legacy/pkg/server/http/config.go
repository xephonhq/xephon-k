package http

import "github.com/pkg/errors"

const (
	DefaultPort = 2333
	MinimalPort = 1024
)

type Config struct {
	Host        string                 `yaml:"host" json:"host"`
	Port        int                    `yaml:"port" json:"port"`
	Enabled     bool                   `yaml:"enabled" json:"enabled"`
	EnablePProf bool                   `yaml:"enablePprof" json:"enablePprof"`
	XXX         map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type configAlias Config

func NewConfig() Config {
	return Config{
		Host:        "localhost",
		Port:        DefaultPort,
		Enabled:     true,
		EnablePProf: false,
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
	// TODO: validate the host
	if c.Port < MinimalPort {
		return errors.Errorf("port number must be greater than %d, got %d instead", MinimalPort, c.Port)
	}
	return nil
}
