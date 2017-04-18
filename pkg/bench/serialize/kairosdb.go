package serialize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xephonhq/xephon-k/pkg/common"
	"io"
	"io/ioutil"
)

type KairosDBSerialize struct {
	buf         bytes.Buffer
	firstSeries bool
}

func (kdb *KairosDBSerialize) Start() {
	kdb.buf.WriteString("[")
	kdb.firstSeries = true
}

func (kdb *KairosDBSerialize) End() {
	kdb.buf.WriteString("]")
}

func (kdb *KairosDBSerialize) Reset() {
	kdb.buf.Reset()
}

func (kdb *KairosDBSerialize) ReadCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(kdb.buf.Bytes()))
}

func (kdb *KairosDBSerialize) Data() []byte {
	return kdb.buf.Bytes()
}

func (kdb *KairosDBSerialize) DataLen() int {
	return kdb.buf.Len()
}

func (kdb *KairosDBSerialize) WriteInt(series common.IntSeries) {
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
