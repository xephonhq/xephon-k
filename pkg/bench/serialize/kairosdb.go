package serialize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xephonhq/xephon-k/pkg/common"
)

type KairosDBSerialize struct {
}

func (kdb *KairosDBSerialize) WriteInt(series common.IntSeries) []byte {
	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprintf("[{\"name\":\"%s\",\"datapoints\":", series.Name))
	j, err := json.Marshal(series.Points)
	if err != nil {
		log.Panicf("can't serialize to json: %s", err.Error())
	}
	buf.Write(j)
	buf.WriteString(",\"tags\":")
	j, err = json.Marshal(series.Tags)
	if err != nil {
		log.Panicf("can't serialize to json: %s", err.Error())
	}
	buf.Write(j)
	buf.WriteString("}]")
	return buf.Bytes()
}
