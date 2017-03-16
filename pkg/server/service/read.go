package service

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/storage"
)

type ReadService interface {
	Service
}

type ReadServiceImpl struct {
	store storage.Store
}

type readRequest struct {
	StartTime int64          `json:"start_time,omitempty"`
	EndTime   int64          `json:"end_time,omitempty"`
	Queries   []common.Query `json:"queries"`
}

type readResponse struct {
	Error    bool                 `json:"error"`
	ErrorMsg string               `json:"error_msg"`
	Queries  []common.QueryResult `json:"queries"` // TODO: maybe we should name it as query result?
}

type ReadServiceHTTPFactory struct {
}

func (ReadServiceHTTPFactory) MakeEndpoint(service Service) endpoint.Endpoint {
	_, ok := service.(ReadService)
	//readSvc, ok := service.(ReadService)
	if !ok {
		log.Panic("must pass read service to read service factory")
	}
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_, ok := request.(readRequest)
		if !ok {
			log.Panic("should be readRequest")
		}
		res := readResponse{}
		return res, nil
	}
}

// TODO: real decode logic
func (ReadServiceHTTPFactory) MakeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		return readRequest{}, nil
	}
}

func (ReadServiceHTTPFactory) MakeEncode() httptransport.EncodeResponseFunc {
	return httptransport.EncodeJSONResponse
}

func NewReadServiceImpl() *ReadServiceImpl {
	return &ReadServiceImpl{}
}

func (ReadServiceImpl) ServiceName() string {
	return "read"
}
