package encoding

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestRawBinaryEncoder_WriteTime(t *testing.T) {
	assert := asst.New(t)

	var (
		encoder TimeEncoder
		b       []byte
	)
	encoder = NewBigEndianBinaryEncoder()
	num := 100
	b = encodeNanoseconds(t, num, encoder)
	// codec + value
	assert.Equal(8*num+1, len(b))

	encoder = NewLittleEndianBinaryEncoder()
	b = encodeNanoseconds(t, num, encoder)
	assert.Equal(8*num+1, len(b))
}

func TestRawBinaryDecoder_ReadTime(t *testing.T) {
	assert := asst.New(t)
	num := 100
	encoded := encodeNanoseconds(t, num, NewBigEndianBinaryEncoder())
	decoder := NewRawBinaryDecoder()
	assert.Nil(decoder.Init(encoded))
	i := 0
	for decoder.Next() {
		i++
		decoder.ReadTime()
	}
	assert.Equal(i, num)
}
