package util

import (
	asst "github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestLogConfig_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)
	valid := `
level: info
source: false
color: false
`
	c := NewLogConfig()
	// FIXED: fatal error: stack overflow, use alias, like filter and read service
	err := yaml.Unmarshal([]byte(valid), &c)
	assert.Nil(err)
	undefinedFields := `
haha: 123
`
	err = yaml.Unmarshal([]byte(undefinedFields), &c)
	//t.Log(err)
	assert.NotNil(err)
	wrongLevel := `
level: haha
`
	err = yaml.Unmarshal([]byte(wrongLevel), &c)
	assert.Nil(err)
	assert.False(c.IsValid())
	assert.NotNil(c.Apply())
}
