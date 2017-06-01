package encoding

import (
	"testing"
	"time"

	asst "github.com/stretchr/testify/assert"
)

// TODO: benchmark for encoding and decoding speed
// TODO: better name and abort when n is too large
func encodeSeconds(t *testing.T, n int, enc TimeEncoder) []byte {
	now := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		// 1s = 10e9 nano seconds
		enc.WriteTime(now + int64(1000000000*i))
	}
	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("cant encode time %v", err)
	}
	return b
}

func TestRawBinaryEncoder_WriteTime(t *testing.T) {
	assert := asst.New(t)

	var (
		encoder TimeEncoder
		b       []byte
	)
	encoder = NewBigEndianBinaryEncoder()
	num := 100
	b = encodeSeconds(t, num, encoder)
	// codec + value
	assert.Equal(8*num+1, len(b))

	encoder = NewLittleEndianBinaryEncoder()
	b = encodeSeconds(t, num, encoder)
	assert.Equal(8*num+1, len(b))
}

func TestRawBinaryDecoder_ReadTime(t *testing.T) {
	assert := asst.New(t)
	num := 100
	encoded := encodeSeconds(t, num, NewBigEndianBinaryEncoder())
	decoder := NewRawBinaryDecoder()
	assert.Nil(decoder.Init(encoded))
	i := 0
	for decoder.Next() {
		i++
		decoder.ReadTime()
	}
	assert.Equal(i, num)
}
