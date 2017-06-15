package influxdb

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

type Serializer struct {
	buf       bytes.Buffer
	prefixBuf bytes.Buffer
}

func NewSerializer() *Serializer {
	s := &Serializer{}
	s.Reset()
	return s
}

func (influx *Serializer) End() {
	// nop
}

func (influx *Serializer) Reset() {
	influx.buf.Reset()
}

func (influx *Serializer) ReadCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(influx.buf.Bytes()))
}

func (influx *Serializer) Data() []byte {
	return influx.buf.Bytes()
}

func (influx *Serializer) DataLen() int {
	return influx.buf.Len()
}

func (influx *Serializer) WriteInt(series common.IntSeries) {
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
