package encoding

import (
	asst "github.com/stretchr/testify/assert"
	"testing"
)

func TestVarIntEncoder_WriteTime(t *testing.T) {
	assert := asst.New(t)
	var (
		encoder TimeEncoder
		b       []byte
	)
	encoder = NewVarIntEncoder()
	num := 100
	b = encodeNanoseconds(t, num, encoder)
	// NOTE: when use nanosecond, it takes 9 bytes to encode what is actually 8 bytes, so it becomes larger
	assert.Equal(9*num+1, len(b))
	// when using second, it is much smaller than max of int64, so it use 5 bytes (<8)
	b = encodeSeconds(t, num, encoder)
	assert.True(len(b) < 9*num+1) // 501
}

func TestVarIntDecoder_ReadTime(t *testing.T) {
	assert := asst.New(t)
	num := 100
	encoded := encodeNanoseconds(t, num, NewVarIntEncoder())
	decoder := NewVarIntDecoder()
	assert.Nil(decoder.Init(encoded))
	i := 0
	for decoder.Next() {
		i++
		decoder.ReadTime()
	}
	assert.Equal(i, num)
}
