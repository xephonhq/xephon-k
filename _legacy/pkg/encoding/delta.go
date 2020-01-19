package encoding

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"math"
)

type DeltaCodecFactory struct {
}

type DeltaEncoder struct {
	b10   []byte
	buf   bytes.Buffer
	codec byte
	vI    int64
	vD    float64
}

type DeltaDecoder struct {
	b   []byte
	cur int
	v   uint64
	vI  int64
	vD  float64
}

func init() {
	registeredCodec = append(registeredCodec, CodecDelta)
	registeredFactory[CodecDelta] = &DeltaCodecFactory{}
}

func NewDeltaEncoder() *DeltaEncoder {
	e := &DeltaEncoder{
		b10:   make([]byte, 10),
		codec: CodecDelta,
	}
	e.Reset()
	return e
}

func NewDeltaDecoder() *DeltaDecoder {
	return &DeltaDecoder{}
}

func (f *DeltaCodecFactory) NewTimeEncoder() (TimeEncoder, error) {
	return NewDeltaEncoder(), nil
}

func (f *DeltaCodecFactory) NewTimeDecoder() (TimeDecoder, error) {
	return NewDeltaDecoder(), nil
}

func (f *DeltaCodecFactory) NewIntValueEncoder() (ValueEncoder, error) {
	return NewDeltaEncoder(), nil
}

func (f *DeltaCodecFactory) NewIntValueDecoder() (ValueDecoder, error) {
	return NewDeltaDecoder(), nil
}

func (f *DeltaCodecFactory) NewDoubleValueEncoder() (ValueEncoder, error) {
	return NewDeltaEncoder(), nil
}

func (f *DeltaCodecFactory) NewDoubleValueDecoder() (ValueDecoder, error) {
	return NewDeltaDecoder(), nil
}

func (e *DeltaEncoder) Codec() byte {
	return e.codec
}

func (e *DeltaEncoder) Bytes() ([]byte, error) {
	return e.buf.Bytes(), nil
}

func (e *DeltaEncoder) Reset() {
	e.buf.Reset()
	e.buf.WriteByte(e.codec)
	e.vI = 0
	e.vD = 0
}

func (e *DeltaEncoder) WriteTime(t int64) {
	// write the delta, we don't need to deal with first value, since the initial value is 0
	n := binary.PutVarint(e.b10, t-e.vI)
	e.buf.Write(e.b10[:n])
	e.vI = t
}

func (e *DeltaEncoder) WriteInt(v int64) {
	n := binary.PutVarint(e.b10, v-e.vI)
	e.buf.Write(e.b10[:n])
	e.vI = v
}

func (e *DeltaEncoder) WriteDouble(v float64) {
	n := binary.PutUvarint(e.b10, math.Float64bits(v-e.vD))
	e.buf.Write(e.b10[:n])
	e.vD = v
}

func (d *DeltaDecoder) Init(b []byte) error {
	// codec + v
	if len(b) < 2 {
		return errors.Wrapf(ErrTooSmall, "got %d but at least 2 bytes is needed codec, value", len(b))
	}
	if b[0] != CodecDelta {
		return errors.Wrapf(ErrCodecMismatch, "DeltaDecoder does not support %s", CodecString(b[0]))
	}
	// exclude codec
	d.b = b[1:]
	d.cur = 0
	// set first value to zero because when we initial the encoder, we set initial value as zero to avoid
	// checking if this is the first value written
	d.vI = 0
	d.vD = 0
	return nil
}

func (d *DeltaDecoder) Next() bool {
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

func (d *DeltaDecoder) ReadTime() int64 {
	// zig zag
	x := int64(d.v >> 1)
	if d.v&1 != 0 {
		x = ^x
	}
	d.vI += x
	return d.vI
}

func (d *DeltaDecoder) ReadInt() int64 {
	// zig zag
	x := int64(d.v >> 1)
	if d.v&1 != 0 {
		x = ^x
	}
	d.vI += x
	return d.vI
}

func (d *DeltaDecoder) ReadDouble() float64 {
	d.vD += math.Float64frombits(d.v)
	return d.vD
}
