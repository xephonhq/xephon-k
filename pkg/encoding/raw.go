package encoding

import (
	"bytes"
	"encoding/binary"
	"math"
)

// Non compression binary encoding using little/big endian

type RawBinaryEncoder struct {
	order binary.ByteOrder
	b8    []byte
	buf   bytes.Buffer
}

func NewBigEndianBinaryEncoder() *RawBinaryEncoder {
	return &RawBinaryEncoder{
		order: binary.BigEndian,
		b8:    make([]byte, 8),
	}
}

func NewLittleEndianBinaryEncoder() *RawBinaryEncoder {
	return &RawBinaryEncoder{
		order: binary.LittleEndian,
		b8:    make([]byte, 8),
	}
}

func (e *RawBinaryEncoder) Codec() byte {
	if e.order.String() == "LittleEndian" {
		return CodecRawLittleEndian
	}
	return CodecRawBigEndian
}

func (e *RawBinaryEncoder) Bytes() ([]byte, error) {
	return e.buf.Bytes(), nil
}

func (e *RawBinaryEncoder) Reset() {
	e.buf.Reset()
}

func (e *RawBinaryEncoder) WriteTime(t int64) {
	e.order.PutUint64(e.b8, uint64(t))
	e.buf.Write(e.b8)
}

func (e *RawBinaryEncoder) WriteInt(v int64) {
	e.order.PutUint64(e.b8, uint64(v))
	e.buf.Write(e.b8)
}

func (e *RawBinaryEncoder) WriteDouble(v float64) {
	e.order.PutUint64(e.b8, math.Float64bits(v))
	e.buf.Write(e.b8)
}
