package encoding

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"math"
)

// Run length encoding without delta, value and length use variable length int encoding
type RLECodecFactory struct {
}

type RLEEncoder struct {
	b10   []byte
	buf   bytes.Buffer
	codec byte
	vI    int64
	vD    float64
	rLen  uint64
}

type RLEDecoder struct {
	b    []byte
	cur  int
	vI   int64
	vD   float64
	rLen uint64
}

func init() {
	registeredCodec = append(registeredCodec, CodecRLE)
	registeredFactory[CodecRLE] = &RLECodecFactory{}
}

func NewRLEEncoder() *RLEEncoder {
	e := &RLEEncoder{
		b10:   make([]byte, 10),
		codec: CodecRLE,
	}
	e.Reset()
	return e
}

func NewRLEDecoder() *RLEDecoder {
	return &RLEDecoder{}
}

func (f *RLECodecFactory) NewTimeEncoder() (TimeEncoder, error) {
	return NewRLEEncoder(), nil
}

func (f *RLECodecFactory) NewTimeDecoder() (TimeDecoder, error) {
	return NewRLEDecoder(), nil
}

func (f *RLECodecFactory) NewIntValueEncoder() (ValueEncoder, error) {
	return NewRLEEncoder(), nil
}

func (f *RLECodecFactory) NewIntValueDecoder() (ValueDecoder, error) {
	return NewRLEDecoder(), nil
}

func (f *RLECodecFactory) NewDoubleValueEncoder() (ValueEncoder, error) {
	return NewRLEEncoder(), nil
}

func (f *RLECodecFactory) NewDoubleValueDecoder() (ValueDecoder, error) {
	return NewRLEDecoder(), nil
}

func (e *RLEEncoder) Codec() byte {
	return e.codec
}

func (e *RLEEncoder) Bytes() ([]byte, error) {
	if e.rLen == 0 {
		return nil, errors.New("run length is 0")
	}
	// write the last run length
	n := binary.PutUvarint(e.b10, e.rLen)
	e.buf.Write(e.b10[:n])
	return e.buf.Bytes(), nil
}

func (e *RLEEncoder) Reset() {
	e.buf.Reset()
	e.buf.WriteByte(e.codec)
	// NOTE: set 0 as our first value and length as 0, otherwise we have to check if it is the first value in all write calls
	e.vI = 0
	e.vD = 0
	e.rLen = 0
}

func (e *RLEEncoder) WriteTime(t int64) {
	if e.vI == t {
		e.rLen++
		return
	}
	// deal with the special value we write at start
	if e.rLen != 0 {
		// write the run length only
		n := binary.PutUvarint(e.b10, e.rLen)
		e.buf.Write(e.b10[:n])
	}

	// the new value
	e.vI = t
	e.rLen = 1
	// write the new value
	n := binary.PutVarint(e.b10, e.vI)
	e.buf.Write(e.b10[:n])
}

func (e *RLEEncoder) WriteInt(v int64) {
	if e.vI == v {
		e.rLen++
		return
	}
	// deal with the special value we write at start
	if e.rLen != 0 {
		// write the run length only
		n := binary.PutUvarint(e.b10, e.rLen)
		e.buf.Write(e.b10[:n])
	}

	// the new value
	e.vI = v
	e.rLen = 1
	// write the new value
	n := binary.PutVarint(e.b10, e.vI)
	e.buf.Write(e.b10[:n])
}

func (e *RLEEncoder) WriteDouble(v float64) {
	if e.vD == v {
		e.rLen++
		return
	}
	if e.rLen != 0 {
		n := binary.PutUvarint(e.b10, e.rLen)
		e.buf.Write(e.b10[:n])
	}

	e.vD = v
	e.rLen = 1
	n := binary.PutUvarint(e.b10, math.Float64bits(v))
	e.buf.Write(e.b10[:n])
}

func (d *RLEDecoder) Init(b []byte) error {
	// codec + v + len
	if len(b) < 3 {
		return errors.Wrapf(ErrTooSmall, "got %d but at least 3 bytes is needed codec, value, length", len(b))
	}
	if b[0] != CodecRLE {
		return errors.Wrapf(ErrCodecMismatch, "RLEDecoder does not support %s", CodecString(b[0]))
	}
	// exclude codec
	d.b = b[1:]
	d.cur = 0
	// read the first value and length
	if err := d.read(); err != nil {
		return errors.Wrap(err, "can't read first value, length pair")
	}
	return nil
}

// NOTE: read also do the convert
func (d *RLEDecoder) read() error {
	v, n := binary.Uvarint(d.b[d.cur:])
	if n <= 0 {
		return errors.New("can't read value")
	}
	d.cur += n
	d.rLen, n = binary.Uvarint(d.b[d.cur:])
	if n <= 0 {
		return errors.New("can't read run length")
	}
	d.cur += n
	// convert
	x := int64(v >> 1)
	if v&1 != 0 {
		x = ^x
	}
	d.vI = x
	d.vD = math.Float64frombits(v)
	return nil
}

func (d *RLEDecoder) Next() bool {
	if d.rLen > 0 {
		d.rLen--
		return true
	}
	if d.cur >= len(d.b) {
		return false
	}
	if err := d.read(); err != nil {
		// TODO: maybe move the cur to the end to avoid calling next again?
		return false
	}
	d.rLen--
	return true
}

func (d *RLEDecoder) ReadTime() int64 {
	return d.vI
}

func (d *RLEDecoder) ReadInt() int64 {
	return d.vI
}

func (d *RLEDecoder) ReadDouble() float64 {
	return d.vD
}
