package xephonk

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"io/ioutil"

	"github.com/xephonhq/xephon-k/pkg/client"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.client.xephonk")

type Client struct {
	config     client.Config
	h          http.Client
	writeReq   *http.Request
	serializer *Serializer
}

func New(config client.Config, transport *http.Transport) (*Client, error) {
	writeReq, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%d/%s", config.Host, config.Port, config.URL), nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		config:     config,
		h:          http.Client{Transport: transport, Timeout: time.Second * time.Duration(config.Timeout)},
		writeReq:   writeReq,
		serializer: &Serializer{},
	}, nil
}

func (c *Client) WriteInt(series *common.IntSeries) {
	c.serializer.WriteInt(*series)
}

func (c *Client) Send() client.Result {
	c.serializer.End()
	result := client.Result{
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
