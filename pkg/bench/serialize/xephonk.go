package serialize

import (
	"bytes"
	"encoding/json"

	"github.com/xephonhq/xephon-k/pkg/common"
)

type XephonKSerialize struct {
}

// WriteInt implements Serializer
func (xk *XephonKSerialize) WriteInt(series common.IntSeries) []byte {
	buf := bytes.NewBufferString("[")
	j, err := json.Marshal(series)
	if err != nil {
		log.Panicf("can't serialize to json: %s", err.Error())
	}
	buf.Write(j)
	buf.WriteString("]")
	return buf.Bytes()
}
