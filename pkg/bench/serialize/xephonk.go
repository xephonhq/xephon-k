package serialize

import (
	"bytes"
	"encoding/json"

	"github.com/xephonhq/xephon-k/pkg/common"
	"io"
	"io/ioutil"
)

type XephonKSerialize struct {
	buf         bytes.Buffer // In most cases, new(Buffer) (or just declaring a Buffer variable) is sufficient to initialize a Buffer.
	firstSeries bool
}

func (xk *XephonKSerialize) Start() {
	xk.buf.WriteString("[")
	xk.firstSeries = true
}

func (xk *XephonKSerialize) End() {
	xk.buf.WriteString("]")
}

func (xk *XephonKSerialize) Reset() {
	// Reset resets the buffer to be empty,
	// but it retains the underlying storage for use by future writes.
	xk.buf.Reset()
}

func (xk *XephonKSerialize) ReadCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(xk.buf.Bytes()))
}

func (xk *XephonKSerialize) Data() []byte {
	return xk.buf.Bytes()
}

func (xk *XephonKSerialize) DataLen() int {
	return xk.buf.Len()
}

// WriteInt implements Serializer
func (xk *XephonKSerialize) WriteInt(series common.IntSeries) {
	if !xk.firstSeries {
		xk.buf.WriteString(",")
	} else {
		xk.firstSeries = false
	}
	j, err := json.Marshal(series)
	if err != nil {
		log.Panicf("can't serialize to json: %s", err.Error())
	}
	xk.buf.Write(j)
}
