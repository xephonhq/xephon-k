package encoding

import (
	asst "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// check interface
var _ TimeEncoder = (*RawBinaryEncoder)(nil)
var _ ValueEncoder = (*RawBinaryEncoder)(nil)
var _ TimeEncoder = (*VarIntEncoder)(nil)
var _ ValueEncoder = (*VarIntEncoder)(nil)

var _ TimeDecoder = (*RawBinaryDecoder)(nil)
var _ ValueDecoder = (*RawBinaryDecoder)(nil)

// TODO: benchmark for encoding and decoding speed
// TODO: better name and abort when n is too large
func encodeNanoseconds(t *testing.T, n int, enc TimeEncoder) []byte {
	enc.Reset()
	now := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		// 1s = 10e9 nano seconds
		enc.WriteTime(now + int64(1000000000*i))
	}
	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("cant encode time %v using %s", err, CodecString(enc.Codec()))
	}
	return b
}

func encodeSeconds(t *testing.T, n int, enc TimeEncoder) []byte {
	enc.Reset()
	now := time.Now().Unix()
	for i := 0; i < n; i++ {
		enc.WriteTime(now + int64(i))
	}
	b, err := enc.Bytes()
	if err != nil {
		t.Fatalf("cant encode time %v using %s", err, CodecString(enc.Codec()))
	}
	return b
}

// TODO: maybe we should do the same for time
func TestRegisteredValueEncoderDecoder(t *testing.T) {
	assert := asst.New(t)
	ivals := []int64{-1, 1}
	dvals := []float64{-1.1, 1.1}

	for i, codec := range registeredCodec {
		assert.NotContains(CodecString(codec), "unkown")
		t.Logf("test %s", CodecString(codec))
		encoder := registeredValueEncoder[i]
		decoder := registeredValueDecoder[i]

		for _, iv := range ivals {
			encoder.WriteInt(iv)
		}
		b, err := encoder.Bytes()
		assert.Nil(err)
		assert.Nil(decoder.Init(b))
		for _, iv := range ivals {
			decoder.Next()
			assert.Equal(iv, decoder.ReadInt())
		}

		encoder.Reset()

		for _, dv := range dvals {
			encoder.WriteDouble(dv)
		}
		b, err = encoder.Bytes()
		assert.Nil(err)
		assert.Nil(decoder.Init(b))
		for _, dv := range dvals {
			decoder.Next()
			assert.Equal(dv, decoder.ReadDouble())
		}
	}
}
