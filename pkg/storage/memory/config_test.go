package memory

import (
	asst "github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestConfig_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)
	valid := `
layout: row
chunkSize: 104857600
enableIndex: true
`
	c := NewConfig()
	err := yaml.Unmarshal([]byte(valid), &c)
	assert.Nil(err)
	undefinedFields := `
haha: 123
`
	err = yaml.Unmarshal([]byte(undefinedFields), &c)
	assert.NotNil(err)
	wrongLayout := `
layout: hybrid`
	err = yaml.Unmarshal([]byte(wrongLayout), &c)
	assert.Nil(err)
	assert.NotNil(c.Validate())
	assert.NotNil(c.Apply())
}
