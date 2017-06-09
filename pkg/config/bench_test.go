package config

import (
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/util"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestBenchConfig_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)
	exampleConfig := util.ReadAsBytes(t, "xkb.yml")
	c := NewBench()
	err := yaml.Unmarshal(exampleConfig, &c)
	assert.NotNil(err)
	assert.Nil(c.Validate())
	assert.Nil(c.Apply())
}
