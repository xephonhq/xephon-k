package encoding

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"math"
)

// NOTE: it is compatible with leb128
type VarIntCodecFactory struct {
}

type VarIntEncoder struct {
	// From encoding/binary At most 10 bytes are needed for 64-bit values
	b10   []byte
	buf   bytes.Buffer
	codec byte
}

type VarIntDecoder struct {
	codec byte
	b     []byte
	cur   int
	v     uint64
}

func init() {
	registeredCodec = append(registeredCodec, CodecVarInt)
	registeredFactory[CodecVarInt] = &VarIntCodecFactory{}
}

func NewVarIntEncoder() *VarIntEncoder {
	e := &VarIntEncoder{
		b10:   make([]byte, 10),
		codec: CodecVarInt,
	}
	e.Reset()
	return e
}

func NewVarIntDecoder() *VarIntDecoder {
	return &VarIntDecoder{}
}

func (f *VarIntCodecFactory) NewTimeEncoder() (TimeEncoder, error) {
	return NewVarIntEncoder(), nil
}

func (f *VarIntCodecFactory) NewTimeDecoder() (TimeDecoder, error) {
	return NewVarIntDecoder(), nil
}

func (f *VarIntCodecFactory) NewIntValueEncoder() (ValueEncoder, error) {
	return NewVarIntEncoder(), nil
}

func (f *VarIntCodecFactory) NewIntValueDecoder() (ValueDecoder, error) {
	return NewVarIntDecoder(), nil
}

func (f *VarIntCodecFactory) NewDoubleValueEncoder() (ValueEncoder, error ){
	return NewVarIntEncoder(), nil
}

func (f *VarIntCodecFactory) NewDoubleValueDecoder() (ValueDecoder, error) {
	return NewVarIntDecoder(), nil
}

func (e *VarIntEncoder) Codec() byte {
	return e.codec
}

func (e *VarIntEncoder) Bytes() ([]byte, error) {
	return e.buf.Bytes(), nil
}

func (e *VarIntEncoder) Reset() {
	e.buf.Reset()
	e.buf.WriteByte(e.codec)
}

func (e *VarIntEncoder) WriteTime(t int64) {
	n := binary.PutVarint(e.b10, t)
	e.buf.Write(e.b10[:n])
}

func (e *VarIntEncoder) WriteInt(v int64) {
	n := binary.PutVarint(e.b10, v)
	e.buf.Write(e.b10[:n])
}

func (e *VarIntEncoder) WriteDouble(v float64) {
	n := binary.PutUvarint(e.b10, math.Float64bits(v))
	e.buf.Write(e.b10[:n])
}

func (d *VarIntDecoder) Init(b []byte) error {
	if len(b) < 2 {
		return errors.Wrap(ErrTooSmall, "at least 2 bytes is needed for codec and a single value")
	}
	if b[0] != CodecVarInt {
		return errors.Wrapf(ErrCodecMismatch, "VarIntDecoder does not support %s", CodecString(b[0]))
	}
	// exclude codec
	d.b = b[1:]
	d.cur = 0
	return nil
}

func (d *VarIntDecoder) Next() bool {
	if d.cur >= len(d.b) {
		return false
	}
	var n int
	d.v, n = binary.Uvarint(d.b[d.cur:])
	if n <= 0 {
		return false
	}
	d.cur += n
	return true
}

func (d *VarIntDecoder) ReadTime() int64 {
	// zig zag
	x := int64(d.v >> 1)
	if d.v&1 != 0 {
		x = ^x
	}
	return x
}

func (d *VarIntDecoder) ReadInt() int64 {
	// zig zag
	x := int64(d.v >> 1)
	if d.v&1 != 0 {
		x = ^x
	}
	return x
}

func (d *VarIntDecoder) ReadDouble() float64 {
	return math.Float64frombits(d.v)
}
