package server

import (
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/server/http"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestConfig_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)
	empty := ``
	c := NewConfig()
	err := yaml.Unmarshal([]byte(empty), &c)
	assert.Nil(err)
	assert.Equal(http.DefaultPort, c.Http.Port)
}
