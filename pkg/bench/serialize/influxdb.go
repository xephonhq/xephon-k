package serialize

import (
	"bytes"
	"fmt"

	"github.com/xephonhq/xephon-k/pkg/common"
	"io"
	"io/ioutil"
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
	buf       bytes.Buffer
	prefixBuf bytes.Buffer
}

func (influx *InfluxDBSerialize) Start() {
	// nop
}

func (influx *InfluxDBSerialize) End() {
	// nop
}

func (influx *InfluxDBSerialize) Reset() {
	influx.buf.Reset()
}

func (influx *InfluxDBSerialize) ReadCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(influx.buf.Bytes()))
}

func (influx *InfluxDBSerialize) Data() []byte {
	return influx.buf.Bytes()
}

func (influx *InfluxDBSerialize) DataLen() int {
	return influx.buf.Len()
}

func (influx *InfluxDBSerialize) WriteInt(series common.IntSeries) {
	// http://herman.asia/efficient-string-concatenation-in-go
	influx.prefixBuf.WriteString(series.Name)
	influx.prefixBuf.WriteString(",")
	tagsLength := len(series.Tags)
	i := 0
	for k, v := range series.Tags {
		i++
		influx.prefixBuf.WriteString(k)
		influx.prefixBuf.WriteString("=")
		influx.prefixBuf.WriteString(v)
		if i < tagsLength {
			influx.prefixBuf.WriteString(",")
		} else {
			influx.prefixBuf.WriteString(" ")
		}
	}
	prefix := influx.prefixBuf.Bytes()
	//log.Info(string(prefix))
	influx.prefixBuf.Reset()

	pointLength := len(series.Points)
	for i = 0; i < pointLength; i++ {
		influx.buf.Write(prefix)
		influx.buf.WriteString(fmt.Sprintf("value=%d %d\n", series.Points[i].V, series.Points[i].T))
	}
}
