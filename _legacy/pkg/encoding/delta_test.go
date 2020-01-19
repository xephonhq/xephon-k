package encoding

import (
	asst "github.com/stretchr/testify/assert"
	"testing"
)

func TestDeltaEncoder_WriteTime(t *testing.T) {
	assert := asst.New(t)
	var (
		encoder TimeEncoder
		b       []byte
	)
	encoder = NewDeltaEncoder()
	num := 100
	b = encodeNanoseconds(t, num, encoder)
	// delta is 1,000,000,000, which takes 5 byte
	// the first value is 9 byte, 1 byte larger than the actual 8 byte
	// 1 byte for codec
	assert.Equal(5*(num-1)+9+1, len(b))
	b = encodeSeconds(t, num, encoder)
	// delta is 1
	// the first value is 5 byte
	assert.Equal(1*(num-1)+5+1, len(b))
}

func TestDeltaDecoder_ReadTime(t *testing.T) {
	assert := asst.New(t)
	num := 100
	encoded := encodeNanoseconds(t, num, NewDeltaEncoder())
	decoder := NewDeltaDecoder()
	assert.Nil(decoder.Init(encoded))
	i := 0
	for decoder.Next() {
		i++
		decoder.ReadTime()
	}
	assert.Equal(i, num)
}
