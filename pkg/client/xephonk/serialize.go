package xephonk

import (
	"bytes"
	"encoding/json"

	"io"
	"io/ioutil"

	"github.com/xephonhq/xephon-k/pkg/common"
)

type Serializer struct {
	buf         bytes.Buffer // In most cases, new(Buffer) (or just declaring a Buffer variable) is sufficient to initialize a Buffer.
	firstSeries bool
}

func NewSeializer() *Serializer {
	s := &Serializer{}
	s.Reset()
	return s
}

func (xk *Serializer) End() {
	xk.buf.WriteString("]")
}

func (xk *Serializer) Reset() {
	// Reset resets the buffer to be empty,
	// but it retains the underlying storage for use by future writes.
	xk.buf.Reset()
	xk.buf.WriteString("[")
	xk.firstSeries = true
}

func (xk *Serializer) ReadCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(xk.buf.Bytes()))
}

func (xk *Serializer) Data() []byte {
	return xk.buf.Bytes()
}

func (xk *Serializer) DataLen() int {
	return xk.buf.Len()
}

// WriteInt implements Serializer
func (xk *Serializer) WriteInt(series common.IntSeries) {
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
