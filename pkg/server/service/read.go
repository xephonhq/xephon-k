package service

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"
)

type ReadService interface {
	Service
	QueryInt(q common.Query) []common.IntSeries
}

type ReadServiceServerImpl struct {
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
	//_, ok := service.(ReadService)
	readSvc, ok := service.(ReadService)
	if !ok {
		log.Panic("must pass read service to read service factory")
	}
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(readRequest)
		if !ok {
			log.Panic("should be readRequest")
		}
		// for all the queries query the data
		results := []common.IntSeries{}
		for _, query := range req.Queries {
			// merge it
			// http://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
			results = append(results, readSvc.QueryInt(query)...)
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

func NewReadServiceMem() *ReadServiceServerImpl {
	store := memory.GetDefaultMemStore()
	return &ReadServiceServerImpl{store: store}
}

func NewReadServiceCassandra(host string) *ReadServiceServerImpl {
	store := cassandra.GetDefaultCassandraStore(host)
	return &ReadServiceServerImpl{store: store}
}

// ServiceName implements Service
func (ReadServiceServerImpl) ServiceName() string {
	return "read"
}

func (rs ReadServiceServerImpl) QueryInt(q common.Query) []common.IntSeries {
	series, err := rs.store.QueryIntSeries(q)
	// TODO: better error handling
	if err != nil {
		log.Warn(err)
	}
	return series
}
