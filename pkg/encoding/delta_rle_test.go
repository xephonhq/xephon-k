package encoding

import (
	asst "github.com/stretchr/testify/assert"
	"testing"
)

func TestDeltaRLEEncoder_WriteTime(t *testing.T) {
	assert := asst.New(t)
	var (
		encoder TimeEncoder
		b       []byte
	)
	encoder = NewDeltaRLEEncoder()
	num := 100
	b = encodeNanoseconds(t, num, encoder)
	//t.Log(b)
	// 1 byte for codec
	// 9 byte for the first timestamp, 1 for its run length
	// 5 byte for 1,000, 000, 000 (nanosecond), num is 100, so length is 99, 1 byte
	assert.Equal(1+9+1+5+1, len(b))
	// [6 202 172 250 165 146 253 152 200 41 1 128 168 214 185 7 99]
	// [6 252 209 137 246 233 253 152 200 41 1 128 168 214 185 7 99]
	b = encodeSeconds(t, num, encoder)
	//t.Log(b)
	// 1 byte for codec
	// 5 byte for first timestamp, 1 for its run length
	// 1 byte for 1,
	// FIXED: why it's 2 instead of 1, length is 99
	// because it is using zig-zag encoding !!!
	// [6 170 214 144 148 11 1 2 99]
	// [6 194 214 144 148 11 1 2 99]
	// [6 208 222 144 148 11 1 2 99]
	assert.Equal(1+5+1+1+1, len(b))
}

func TestDeltaRLEDecoder_ReadTime(t *testing.T) {
	assert := asst.New(t)
	num := 100
	encoded := encodeNanoseconds(t, num, NewDeltaRLEEncoder())
	decoder := NewDeltaRLEDecoder()
	assert.Nil(decoder.Init(encoded))
	i := 0
	for decoder.Next() {
		i++
		decoder.ReadTime()
	}
	assert.Equal(i, num)
}
