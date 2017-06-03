package encoding

import (
	"errors"
	"fmt"

	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.encoding")

const (
	_                    byte = iota
	CodecRawBigEndian
	CodecRawLittleEndian
	CodecVarInt
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
	// TODO: Once or Batch?
	//WriteTimeOnce(t []int64)
}

type TimeDecoder interface {
	Init([]byte) error
	Next() bool
	ReadTime() int64
	// TODO: ReadAll
	// TODO: Reset
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

// TODO: test we support all the codec's, maybe have a registered codec
func CodecString(codec byte) string {
	switch codec {
	case CodecRawBigEndian:
		return "codec: raw bigendian"
	case CodecRawLittleEndian:
		return "codec: raw littleendian"
	case CodecVarInt:
		return "codec: variable length integer"
	default:
		return fmt.Sprintf("codec: unknown %d", codec)
	}
}
