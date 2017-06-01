package encoding

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/pkg/errors"
)

// Non compression binary encoding using little/big endian

type RawBinaryEncoder struct {
	order binary.ByteOrder
	b8    []byte
	buf   bytes.Buffer
	codec byte
}

type RawBinaryDecoder struct {
	order binary.ByteOrder
	codec byte
	b     []byte
	cur   int
	v     uint64
}

func NewBigEndianBinaryEncoder() *RawBinaryEncoder {
	e := &RawBinaryEncoder{
		order: binary.BigEndian,
		b8:    make([]byte, 8),
		codec: CodecRawBigEndian,
	}
	e.Reset()
	return e
}

func NewLittleEndianBinaryEncoder() *RawBinaryEncoder {
	e := &RawBinaryEncoder{
		order: binary.LittleEndian,
		b8:    make([]byte, 8),
		codec: CodecRawLittleEndian,
	}
	e.Reset()
	return e
}

// TODO: might split it into two functions, little and big endian
// then the problem is do we still pass the codec byte to Init and set the endian
func NewRawBinaryDecoder() *RawBinaryDecoder {
	// it seems we don't have anything to initialize, the endian is set upon Init
	return &RawBinaryDecoder{}
}

func (e *RawBinaryEncoder) Codec() byte {
	return e.codec
}

func (e *RawBinaryEncoder) Bytes() ([]byte, error) {
	return e.buf.Bytes(), nil
}

func (e *RawBinaryEncoder) Reset() {
	e.buf.Reset()
	e.buf.WriteByte(e.codec)
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

func (d *RawBinaryDecoder) Init(b []byte) error {
	// TODO: if we use gzip, then this should be different, but the gzip wrapper should pass the data after unzip
	if len(b) < 9 {
		return errors.Wrap(ErrTooSmall, "at least 9 bytes is needed for codec and a single value")
	}
	if (len(b)-1)%8 != 0 {
		return errors.Wrap(ErrCodecMismatch, "raw binary encoding would always haves bytes length multiply of 8 when codec is excluded")
	}
	switch b[0] {
	case CodecRawBigEndian:
		d.codec = CodecRawBigEndian
		d.order = binary.BigEndian
	case CodecRawLittleEndian:
		d.codec = CodecRawLittleEndian
		d.order = binary.LittleEndian
	default:
		return errors.Wrapf(ErrCodecMismatch, "RawBinaryDecoder only supports raw binary little/big endian but got %s", CodecString(b[0]))
	}
	// exclude the codec
	d.b = b[1:]
	return nil
}

func (d *RawBinaryDecoder) Next() bool {
	prev := d.cur
	d.cur += 8
	if d.cur > len(d.b) {
		return false
	}
	d.v = d.order.Uint64(d.b[prev:d.cur])
	return true
}

func (d *RawBinaryDecoder) ReadTime() int64 {
	return int64(d.v)
}

func (d *RawBinaryDecoder) ReadInt() int64 {
	return int64(d.v)
}

func (d *RawBinaryDecoder) ReadDouble() float64 {
	return math.Float64frombits(d.v)
}
