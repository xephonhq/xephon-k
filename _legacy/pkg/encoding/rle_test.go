package encoding

import (
	asst "github.com/stretchr/testify/assert"

	"testing"
)

func TestRLEEncoder_WriteTime(t *testing.T) {
	assert := asst.New(t)
	var (
		encoder TimeEncoder
		b       []byte
	)
	encoder = NewRLEEncoder()
	num := 100
	b = encodeNanoseconds(t, num, encoder)
	// NOTE: 9 byte for value, 1 for length, without delta, RLE is pretty useless for encoding regular interval time
	assert.Equal(10*num+1, len(b))
}

func TestRLEDecoder_ReadTime(t *testing.T) {
	assert := asst.New(t)
	num := 100
	encoded := encodeNanoseconds(t, num, NewRLEEncoder())
	decoder := NewRLEDecoder()
	assert.Nil(decoder.Init(encoded))
	i := 0
	for decoder.Next() {
		i++
		decoder.ReadTime()
	}
	assert.Equal(i, num)
}
