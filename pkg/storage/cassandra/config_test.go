package cassandra

import (
	asst "github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestConfig_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)
	valid := `
host: localhost
port: 9042
`
	c := NewConfig()
	err := yaml.Unmarshal([]byte(valid), &c)
	assert.Nil(err)
	undefinedFields := `
haha: 123
`
	err = yaml.Unmarshal([]byte(undefinedFields), &c)
	assert.NotNil(err)
	wrongPort := `
port: -1
`
	err = yaml.Unmarshal([]byte(wrongPort), &c)
	assert.Nil(err)
	assert.NotNil(c.Validate())
	assert.NotNil(c.Apply())
}
