package encoding

import (
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.encoding")

const (
	_ byte = iota
	CodecRawBigEndian
	CodecRawLittleEndian
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
	Init([]byte)
}

type ValueEncoder interface {
	Codec() byte
	Bytes() ([]byte, error)
	Reset()
	WriteInt(v int64)
	WriteDouble(v float64)
}

type ValueDecoder interface {
	Init([]byte)
}

// check interface
var _ TimeEncoder = (*RawBinaryEncoder)(nil)
var _ ValueEncoder = (*RawBinaryEncoder)(nil)
