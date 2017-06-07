package disk

import (
	asst "github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestConfig_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)
	valid := `
folder: /tmp
concurrentWriteFiles: 1
singleFileSize: 536870912
`
	c := NewConfig()
	err := yaml.Unmarshal([]byte(valid), &c)
	assert.Nil(err)
	undefinedFields := `
haha: 123
`
	err = yaml.Unmarshal([]byte(undefinedFields), &c)
	assert.NotNil(err)
	wrongConcurrentWrieFiles := `
concurrentWriteFiles: 2
`
	err = yaml.Unmarshal([]byte(wrongConcurrentWrieFiles), &c)
	assert.Nil(err)
	assert.NotNil(c.Validate())
	assert.NotNil(c.Apply())
	nonExistFolder := `
folder: /whooo/hahaha
`
	err = yaml.Unmarshal([]byte(nonExistFolder), &c)
	assert.Nil(err)
	assert.NotNil(c.Validate())
	assert.NotNil(c.Apply())
}
