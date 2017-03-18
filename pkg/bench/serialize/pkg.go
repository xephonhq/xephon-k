package serialize

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"

	"time"
)

var log = util.Logger.NewEntryWithPkg("k.b.loader")

// check interface
var _ Serializer = (*InfluxDBSerialize)(nil)

// Serializer truns a single series into the bytes payload for certain backend
type Serializer interface {
	WriteInt(common.IntSeries) []byte
	//WritDouble(common.DoubleSeries) []byte
}

// util function for test
func createDummyIntPoints() common.IntSeries {
	tags := make(map[string]string)
	tags["os"] = "ubuntu"
	tags["machine"] = "machine-01"
	startTime := time.Now().Unix() * 1000
	s := common.IntSeries{Name:"cpi", Tags:tags}
	for i := 0; i < 5; i++ {
		s.Points = append(s.Points, common.IntPoint{TimeNano: startTime + int64(i*1000), V: i})
	}
	return s
}
