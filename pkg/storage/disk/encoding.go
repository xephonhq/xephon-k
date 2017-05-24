package disk

import (
	"bytes"
	"encoding/binary"
)

const (
	CompressionNone = iota
	CompressionGzip
	CompressionZlib
)

const (
	EncodingTimeRawBigEndian byte = 1
	EncodingIntRawBigEndian  byte = 2
)

// check interface
var _ TimeEncoder = (*RawBigEndianTimeEncoder)(nil)
var _ IntEncoder = (*RawBigEndianIntEncoder)(nil)

type TimeEncoder interface {
	Encoding() byte
	Write(t int64)
	// TODO: may need to return header information, or handle it in Reset
	Bytes() ([]byte, error)
	Reset()
}

type ValueEncoder interface {
	Encoding() byte
	Bytes() ([]byte, error)
	Reset()
}

type IntEncoder interface {
	ValueEncoder
	Write(v int64)
}

type TimeDecoder interface {
}

type RawBigEndianTimeEncoder struct {
	buf bytes.Buffer
	b   [8]byte
}

// TODO: this is 100% same as RawBigEndianTimeEncoder
type RawBigEndianIntEncoder struct {
	buf bytes.Buffer
	b   [8]byte
}

func (e *RawBigEndianTimeEncoder) Encoding() byte {
	return EncodingTimeRawBigEndian
}

func (e *RawBigEndianTimeEncoder) Write(t int64) {
	// FIXME: why there is not PutInt64
	binary.BigEndian.PutUint64(e.b[:], uint64(t))
	// TODO: should I track error here
	e.buf.Write(e.b[:])
}

func (e *RawBigEndianTimeEncoder) Bytes() ([]byte, error) {
	return e.buf.Bytes(), nil
}

func (e *RawBigEndianTimeEncoder) Reset() {
	e.buf.Reset()
	// TODO: may need to pre fill header information, like for delta/RLE encoding
}

func (e *RawBigEndianIntEncoder) Encoding() byte {
	return EncodingIntRawBigEndian
}

func (e *RawBigEndianIntEncoder) Write(t int64) {
	// FIXME: why there is not PutInt64
	binary.BigEndian.PutUint64(e.b[:], uint64(t))
	// TODO: should I track error here
	e.buf.Write(e.b[:])
}

func (e *RawBigEndianIntEncoder) Bytes() ([]byte, error) {
	return e.buf.Bytes(), nil
}

func (e *RawBigEndianIntEncoder) Reset() {
	e.buf.Reset()
	// TODO: may need to pre fill header information, like for delta/RLE encoding
}
