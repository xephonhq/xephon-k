package encoding

import (
	"errors"
	"fmt"

	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.encoding")

const (
	_ byte = iota
	CodecRawBigEndian
	CodecRawLittleEndian
)

var (
	ErrTooSmall      = errors.New("data for decoding is too small")
	ErrCodecMismatch = errors.New("decoder got data encoded using other codec")
)

type TimeEncoder interface {
	Codec() byte
	Bytes() ([]byte, error)
	Reset()
	WriteTime(t int64)
	// TODO Once or Batch?
	//WriteTimeOnce(t []int64)
}

type TimeDecoder interface {
	Init([]byte) error
	Next() bool
	ReadTime() int64
}

type ValueEncoder interface {
	Codec() byte
	Bytes() ([]byte, error)
	Reset()
	WriteInt(v int64)
	WriteDouble(v float64)
}

type ValueDecoder interface {
	Init([]byte) error
	Next() bool
	ReadInt() int64
	ReadDouble() float64
}

// check interface
var _ TimeEncoder = (*RawBinaryEncoder)(nil)
var _ ValueEncoder = (*RawBinaryEncoder)(nil)

//var _ TimeDecoder = (*RawBinaryDecoder)(nil)
//var _ ValueDecoder = (*RawBinaryDecoder)(nil)

func CodecString(codec byte) string {
	switch codec {
	case CodecRawBigEndian:
		return "codec: raw bigendian"
	case CodecRawLittleEndian:
		return "codec: raw littleendian"
	default:
		return fmt.Sprintf("codec: unknown %d", codec)
	}
}
