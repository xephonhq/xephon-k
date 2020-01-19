package memory

import (
	"github.com/pkg/errors"
)

const (
	MinimalChunkSize = 10 * 1024 * 1024
)

type Config struct {
	Layout      string                 `yaml:"layout" json:"layout"`
	ChunkSize   int                    `yaml:"chunkSize" json:"chunkSize"`
	EnableIndex bool                   `yaml:"enableIndex" json:"enableIndex"`
	XXX         map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type configAlias Config

func NewConfig() Config {
	return Config{
		Layout:      "row",
		ChunkSize:   MinimalChunkSize,
		EnableIndex: true,
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
	// TODO: do we really need to shrink chunk or change the layout?
	// this requires the config have access to the storage, which is inverse of the normal flow
	return nil
}

func (c *Config) Validate() error {
	if c.Layout != "row" && c.Layout != "column" {
		return errors.Errorf("unsupported memory layout %s", c.Layout)
	}
	if c.ChunkSize < MinimalChunkSize {
		return errors.Errorf("chunk size must be larger than 10MB, got %d bytes", c.ChunkSize)
	}
	return nil
}
