package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/xephonhq/xephon-k/pkg/common"
)

type GenericHTTPClient struct {
	config     Config
	h          http.Client
	writeReq   *http.Request
	serializer Serializer
}

func New(config Config, transport *http.Transport, serializer Serializer) (*GenericHTTPClient, error) {
	writeReq, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%d/%s", config.Host, config.Port, config.URL), nil)
	if err != nil {
		return nil, err
	}
	serializer.Reset()
	return &GenericHTTPClient{
		config:     config,
		h:          http.Client{Transport: transport, Timeout: time.Second * time.Duration(config.Timeout)},
		writeReq:   writeReq,
		serializer: serializer,
	}, nil
}

func (c *GenericHTTPClient) WriteInt(series *common.IntSeries) {
	c.serializer.WriteInt(*series)
}

func (c *GenericHTTPClient) Send() Result {
	c.serializer.End()
	result := Result{
		Start:        time.Now(),
		RequestSize:  int64(c.serializer.DataLen()),
		ResponseSize: 0,
	}
	req := new(http.Request)
	*req = *c.writeReq
	req.Body = c.serializer.ReadCloser()
	res, err := c.h.Do(req)
	c.serializer.Reset()
	if err != nil {
		log.Warn(err)
	} else {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}
	if res != nil {
		result.Code = res.StatusCode
		result.ResponseSize = res.ContentLength
	}
	result.Err = err
	result.End = time.Now()
	return result
}
