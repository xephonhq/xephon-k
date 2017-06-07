package server

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/server/grpc"
	"github.com/xephonhq/xephon-k/pkg/server/http"
)

type Config struct {
	Http http.Config            `yaml:"http" json:"http"`
	Grpc grpc.Config            `yaml:"grpc" json:"grpc"`
	XXX  map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type configAlias Config

func NewConfig() Config {
	return Config{
		Http: http.NewConfig(),
		Grpc: grpc.NewConfig(),
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
	if err := c.Http.Validate(); err != nil {
		return err
	}
	if err := c.Grpc.Validate(); err != nil {
		return err
	}
	if c.Http.Port == c.Grpc.Port {
		// FIXME: I remember go-kit can run http and grpc on same port
		return errors.New("can't run http and grpc service on same port")
	}
	return nil
}
