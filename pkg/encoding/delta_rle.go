package encoding

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"math"
)

type DeltaRLECodecFactory struct {
}

type DeltaRLEEncoder struct {
	b10   []byte
	buf   bytes.Buffer
	codec byte
	vI    int64
	dI    int64
	vD    float64
	dD    float64
	rLen  uint64
}

type DeltaRLEDecoder struct {
	b    []byte
	cur  int
	vI   int64
	dI   int64
	vD   float64
	dD   float64
	rLen uint64
}

func init() {
	registeredCodec = append(registeredCodec, CodecDeltaRLE)
	registeredFactory[CodecDeltaRLE] = &DeltaCodecFactory{}
}

func NewDeltaRLEEncoder() *DeltaRLEEncoder {
	e := &DeltaRLEEncoder{
		b10:   make([]byte, 10),
		codec: CodecDeltaRLE,
	}
	e.Reset()
	return e
}

func NewDeltaRLEDecoder() *DeltaRLEDecoder {
	return &DeltaRLEDecoder{}
}

func (f *DeltaRLECodecFactory) NewTimeEncoder() (TimeEncoder, error) {
	return NewDeltaRLEEncoder(), nil
}

func (f *DeltaRLECodecFactory) NewTimeDecoder() (TimeDecoder, error) {
	return NewDeltaRLEDecoder(), nil
}

func (f *DeltaRLECodecFactory) NewIntValueEncoder() (ValueEncoder, error) {
	return NewDeltaRLEEncoder(), nil
}

func (f *DeltaRLECodecFactory) NewIntValueDecoder() (ValueDecoder, error) {
	return NewDeltaRLEDecoder(), nil
}

func (f *DeltaRLECodecFactory) NewDoubleValueEncoder() (ValueEncoder, error) {
	return NewDeltaRLEEncoder(), nil
}

func (f *DeltaRLECodecFactory) NewDoubleValueDecoder() (ValueDecoder, error) {
	return NewDeltaRLEDecoder(), nil
}

func (e *DeltaRLEEncoder) Codec() byte {
	return e.codec
}

func (e *DeltaRLEEncoder) Bytes() ([]byte, error) {
	if e.rLen == 0 {
		return nil, errors.New("run length is 0")
	}
	// write the last run length
	n := binary.PutUvarint(e.b10, e.rLen)
	e.buf.Write(e.b10[:n])
	return e.buf.Bytes(), nil
}

func (e *DeltaRLEEncoder) Reset() {
	e.buf.Reset()
	e.buf.WriteByte(e.codec)
	e.vI = 0
	e.dI = 0
	e.vD = 0
	e.dD = 0
	e.rLen = 0
}

func (e *DeltaRLEEncoder) WriteTime(t int64) {
	d := t - e.vI
	e.vI = t
	if e.dI == d {
		e.rLen++
		return
	}
	// deal with the special value we write at start
	if e.rLen != 0 {
		// write the run length only
		n := binary.PutUvarint(e.b10, e.rLen)
		e.buf.Write(e.b10[:n])
	}
	// the new delta
	e.dI = d
	e.rLen = 1
	// write the new delta
	n := binary.PutVarint(e.b10, e.dI)
	e.buf.Write(e.b10[:n])
}

// TODO: merge WriteTime with WriteInt, this copy and paste in all the encoding code is really xx
func (e *DeltaRLEEncoder) WriteInt(v int64) {
	d := v - e.vI
	e.vI = v
	if e.dI == d {
		e.rLen++
		return
	}
	// deal with the special value we write at start
	if e.rLen != 0 {
		// write the run length only
		n := binary.PutUvarint(e.b10, e.rLen)
		e.buf.Write(e.b10[:n])
	}
	// the new delta
	e.dI = d
	e.rLen = 1
	// write the new delta
	n := binary.PutVarint(e.b10, e.dI)
	e.buf.Write(e.b10[:n])
}

func (e *DeltaRLEEncoder) WriteDouble(v float64) {
	d := v - e.vD
	e.vD = v
	if e.dD == d {
		e.rLen++
		return
	}
	if e.rLen != 0 {
		n := binary.PutUvarint(e.b10, e.rLen)
		e.buf.Write(e.b10[:n])
	}
	e.dD = d
	e.rLen = 1
	n := binary.PutUvarint(e.b10, math.Float64bits(e.dD))
	e.buf.Write(e.b10[:n])
}

func (d *DeltaRLEDecoder) Init(b []byte) error {
	// codec + v + len
	if len(b) < 3 {
		return errors.Wrapf(ErrTooSmall, "got %d but at least 3 bytes is needed codec, value, length", len(b))
	}
	if b[0] != CodecDeltaRLE {
		return errors.Wrapf(ErrCodecMismatch, "DeltaRLEDecoder does not support %s", CodecString(b[0]))
	}
	// exclude codec
	d.b = b[1:]
	d.cur = 0
	d.vI = 0
	d.vD = 0
	// read the first
	if err := d.read(); err != nil {
		return errors.Wrap(err, "can't read first value, length pair")
	}
	return nil
}

func (d *DeltaRLEDecoder) read() error {
	v, n := binary.Uvarint(d.b[d.cur:])
	if n <= 0 {
		return errors.New("can't read value")
	}
	d.cur += n
	d.rLen, n = binary.Uvarint(d.b[d.cur:])
	d.cur += n
	// convert
	x := int64(v >> 1)
	if v&1 != 0 {
		x = ^x
	}
	d.dI = x
	d.dD = math.Float64frombits(v)
	return nil
}

func (d *DeltaRLEDecoder) Next() bool {
	if d.rLen > 0 {
		d.rLen--
		return true
	}
	if d.cur >= len(d.b) {
		return false
	}
	if err := d.read(); err != nil {
		return false
	}
	d.rLen--
	return true
}

func (d *DeltaRLEDecoder) ReadTime() int64 {
	d.vI += d.dI
	return d.vI
}

func (d *DeltaRLEDecoder) ReadInt() int64 {
	d.vI += d.dI
	return d.vI
}

func (d *DeltaRLEDecoder) ReadDouble() float64 {
	d.vD += d.dD
	return d.vD
}
