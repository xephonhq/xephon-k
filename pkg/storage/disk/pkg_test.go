package disk

import (
	"bytes"
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestMagicBytes(t *testing.T) {
	assert := asst.New(t)
	assert.True(IsMagic(MagicBytes()))
}

func TestIsMagic(t *testing.T) {
	assert := asst.New(t)
	var buf bytes.Buffer
	buf.WriteString("xephon-k")
	assert.True(IsMagic(buf.Bytes()))
	assert.False(IsMagic([]byte("x")))
}
