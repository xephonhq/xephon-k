package kairosdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xephonhq/xephon-k/pkg/common"
	"io"
	"io/ioutil"
)

type Serializer struct {
	buf         bytes.Buffer
	firstSeries bool
}

func NewSerializer() *Serializer {
	s := &Serializer{}
	s.Reset()
	return s
}

func (kdb *Serializer) End() {
	kdb.buf.WriteString("]")
}

func (kdb *Serializer) Reset() {
	kdb.buf.Reset()
	kdb.buf.WriteString("[")
	kdb.firstSeries = true
}

func (kdb *Serializer) ReadCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(kdb.buf.Bytes()))
}

func (kdb *Serializer) Data() []byte {
	return kdb.buf.Bytes()
}

func (kdb *Serializer) DataLen() int {
	return kdb.buf.Len()
}

func (kdb *Serializer) WriteInt(series common.IntSeries) {
	if !kdb.firstSeries {
		kdb.buf.WriteString(",")
	} else {
		kdb.firstSeries = false
	}
	kdb.buf.WriteString(fmt.Sprintf("{\"name\":\"%s\",\"datapoints\":", series.Name))
	j, err := json.Marshal(series.Points)
	if err != nil {
		log.Panicf("can't serialize to json: %s", err.Error())
	}
	kdb.buf.Write(j)
	kdb.buf.WriteString(",\"tags\":")
	j, err = json.Marshal(series.Tags)
	if err != nil {
		log.Panicf("can't serialize to json: %s", err.Error())
	}
	kdb.buf.Write(j)
	kdb.buf.WriteString("}")
}
