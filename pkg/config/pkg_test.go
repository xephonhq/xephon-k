package config

import (
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/util"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestDaemonConfig_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)
	exampleConfig := util.ReadAsBytes(t, "xkd.yml")
	c := NewDaemon()
	err := yaml.Unmarshal(exampleConfig, &c)
	assert.Nil(err)
	assert.Nil(c.Validate())
	assert.Nil(c.Apply())
}
