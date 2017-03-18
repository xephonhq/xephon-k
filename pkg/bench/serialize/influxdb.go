package serialize

import (
	"bytes"
	"fmt"

	"github.com/xephonhq/xephon-k/pkg/common"
)

/*
https://docs.influxdata.com/influxdb/v1.2/write_protocols/line_protocol_tutorial/
weather,location=us-midwest temperature=82 1465839830100400200
|    -------------------- --------------  |
|             |             |             |
|             |             |             |
+-----------+--------+-+---------+-+---------+
|measurement|,tag_set| |field_set| |timestamp|
+-----------+--------+-+---------+-+---------+
// localhost:8086/write?db=sb
*/

type InfluxDBSerialize struct {
}

func (influx *InfluxDBSerialize) WriteInt(series common.IntSeries) []byte {
	// http://herman.asia/efficient-string-concatenation-in-go
	buf := bytes.NewBufferString("")
	buf.WriteString(series.Name)
	buf.WriteString(",")
	tagsLength := len(series.Tags)
	i := 0
	for k, v := range series.Tags {
		i++
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		if i < tagsLength {
			buf.WriteString(",")
		} else {
			buf.WriteString(" ")
		}
	}
	prefix := buf.Bytes()
	//log.Info(string(prefix))
	buf.Reset()
	pointLength := len(series.Points)
	for i = 0; i < pointLength; i++ {
		buf.Write(prefix)
		buf.WriteString(fmt.Sprintf("value=%d %d\n", series.Points[i].V, series.Points[i].TimeNano))
	}
	// TODO: recycle the buffer
	return buf.Bytes()
}
