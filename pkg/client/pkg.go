package client

import (
	"time"

	"github.com/xephonhq/xephon-k/pkg/common"
)

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
