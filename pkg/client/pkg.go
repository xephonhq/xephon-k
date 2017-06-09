package client

import (
	"io"
	"time"

	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.client")

type Result struct {
	Err          error
	Code         int
	Start        time.Time
	End          time.Time
	RequestSize  int64
	ResponseSize int64
}

type TSDBClient interface {
	WriteInt(series *common.IntSeries)
	Send() Result
}

// Serializer turns a single series into the bytes payload for certain backend
type Serializer interface {
	End()
	Reset()
	ReadCloser() io.ReadCloser
	Data() []byte
	DataLen() int
	WriteInt(common.IntSeries)
	//WritDouble(common.DoubleSeries) []byte
}
