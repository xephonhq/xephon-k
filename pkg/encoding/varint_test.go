package encoding

import (
	"testing"
	asst "github.com/stretchr/testify/assert"
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

// TODO: use table test and merge with other decoder
func TestVarIntDecoder_ReadInt(t *testing.T) {
	assert := asst.New(t)
	encoder := NewVarIntEncoder()
	encoder.WriteInt(-1)
	encoder.WriteInt(1)
	decoder := NewVarIntDecoder()
	p, err := encoder.Bytes()
	//t.Log(p)
	assert.Nil(err)
	assert.Nil(decoder.Init(p))
	decoder.Next()
	assert.Equal(int64(-1), decoder.ReadInt())
	decoder.Next()
	assert.Equal(int64(1), decoder.ReadInt())
}

// TODO: use table test and merge with other decoder
func TestVarIntDecoder_ReadDouble(t *testing.T) {
	assert := asst.New(t)
	encoder := NewVarIntEncoder()
	encoder.WriteDouble(-1.1)
	encoder.WriteDouble(1.1)
	decoder := NewVarIntDecoder()
	p, err := encoder.Bytes()
	//t.Log(p)
	assert.Nil(err)
	assert.Nil(decoder.Init(p))
	decoder.Next()
	assert.Equal(-1.1, decoder.ReadDouble())
	decoder.Next()
	assert.Equal(1.1, decoder.ReadDouble())
}
