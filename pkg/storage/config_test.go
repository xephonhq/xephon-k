package storage

import (
	asst "github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestConfig_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)
	valid := `
memory:
    layout: row
    chunkSize: 104857600
    enableIndex: true
disk:
    folder: /tmp
    concurrentWriteFiles: 1
    singleFileSize: 536870912
cassandra:
    host: localhost
    port: 9042
`
	c := NewConfig()
	err := yaml.Unmarshal([]byte(valid), &c)
	assert.Nil(err)
	assert.Nil(c.Validate())
	assert.Nil(c.Apply())
	undefinedFields := `
haha: 123
`
	err = yaml.Unmarshal([]byte(undefinedFields), &c)
	assert.NotNil(err)
	wrongLayout := `
memory:
    layout: hybrid
`
	err = yaml.Unmarshal([]byte(wrongLayout), &c)
	assert.Nil(err)
	assert.NotNil(c.Validate())
	assert.NotNil(c.Apply())
}
