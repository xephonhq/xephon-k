package disk

import (
	"bytes"
	asst "github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMagic(t *testing.T) {
	assert := asst.New(t)
	var buf bytes.Buffer
	buf.WriteString("xephon-k")
	assert.True(IsMagic(buf.Bytes()))
	assert.False(IsMagic([]byte("x")))
}
