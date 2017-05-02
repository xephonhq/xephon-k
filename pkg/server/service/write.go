package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"
)

type WriteService interface {
	Service
	WriteInt([]common.IntSeries) error
}

// WriteServiceServerImpl is the server implementation of WriteService
type WriteServiceServerImpl struct {
	store storage.Store
}

type writeRequest struct {
	series []common.IntSeries
}

// TODO: actually we could also tell the client about how many points are written and the duplication etc.
type writeResponse struct {
	Error    bool   `json:"error"`
	ErrorMsg string `json:"error_msg"`
}

// WriteServiceHTTPFactory is used to create the endpoint, encode, decode
type WriteServiceHTTPFactory struct {
}

func (WriteServiceHTTPFactory) MakeEndpoint(service Service) endpoint.Endpoint {
	writeSvc, ok := service.(WriteService)
	if !ok {
		log.Panic("must pass write service to write service factory")
	}
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(writeRequest)
		if !ok {
			log.Panic("should be writeRequest")
		}
		res := writeResponse{Error: false, ErrorMsg: ""}
		err := writeSvc.WriteInt(req.series)
		if err != nil {
			res.Error = true
			res.ErrorMsg = err.Error()
			return res, nil
		}
		return res, nil
	}
}

func (WriteServiceHTTPFactory) MakeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		// NOTE: do I need to allocate space? no
		var series []common.IntSeries
		// FIXME: go-kit does not handle decode error?
		// https://github.com/xephonhq/xephon-k/issues/6
		// https://github.com/go-kit/kit/issues/133
		if err := json.NewDecoder(r.Body).Decode(&series); err != nil {
			return nil, err
		}
		log.Tracef("got %d series after decode ", len(series))
		return writeRequest{series: series}, nil
	}
}

func (WriteServiceHTTPFactory) MakeEncode() httptransport.EncodeResponseFunc {
	return httptransport.EncodeJSONResponse
}

func NewWriteServiceMem() *WriteServiceServerImpl {
	// FIXME: it should be a singleton
	store := memory.GetDefaultMemStore()
	return &WriteServiceServerImpl{store: store}
}

func NewWriteServiceCassandra(host string) *WriteServiceServerImpl {
	store := cassandra.GetDefaultCassandraStore(host)
	return &WriteServiceServerImpl{store: store}
}

// ServiceName implements Service
func (WriteServiceServerImpl) ServiceName() string {
	return "write"
}

// WriteInt implements WriteService
func (ws WriteServiceServerImpl) WriteInt(series []common.IntSeries) error {
	// write to memory storage
	// NOTE: maybe we should wrap error instead of just return it
	return ws.store.WriteIntSeries(series)
}
