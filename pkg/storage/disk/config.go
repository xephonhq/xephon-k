package disk

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

const (
	MinimalSingleFileSize = 64 * 1024 * 1024
)

// TODO: encoding
type Config struct {
	Folder               string                 `yaml:"folder" json:"folder"`
	ConcurrentWriteFiles int                    `yaml:"concurrentWriteFiles" json:"concurrentWriteFiles"`
	SingleFileSize       int                    `yaml:"singleFileSize" json:"singleFileSize"`
	XXX                  map[string]interface{} `yaml:",inline"`
}

// avoid recursion in UnmarshalYAML
type configAlias Config

func NewConfig() Config {
	return Config{
		Folder:               "/tmp",
		ConcurrentWriteFiles: 1,
		SingleFileSize:       MinimalSingleFileSize,
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
	f, err := ioutil.TempFile(c.Folder, "xephonk-diskconfig-probe")
	if err != nil {
		return errors.Wrap(err, "can't write file in specified folder")
	}
	f.Close()
	os.Remove(f.Name())

	if c.ConcurrentWriteFiles != 1 {
		return errors.Errorf("only support write to single file, but got %d", c.ConcurrentWriteFiles)
	}
	if c.SingleFileSize < MinimalSingleFileSize {
		return errors.Errorf("single file size must be larger than 64MB, got %d bytes", c.SingleFileSize)
	}
	return nil
}
