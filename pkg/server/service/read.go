package service

import (
	"context"
	"net/http"

	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
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
	// TODO: where is the data?
}

type ReadServiceHTTPFactory struct {
}

func (ReadServiceHTTPFactory) MakeEndpoint(service Service) endpoint.Endpoint {
	// TODO: test it
	readSvc, ok := service.(ReadService)
	if !ok {
		log.Panic("must pass read service to read service factory")
	}
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(readRequest)
		if !ok {
			log.Panic("should be readRequest")
		}
		res := readResponse{}
		// TODO: check start end time and return 400
		if req.StartTime == 0 || req.EndTime == 0 {
			return res, errors.New("must set start and end time")
		}
		// for all the queries query the data
		results := []common.IntSeries{}
		for _, query := range req.Queries {
			// TODO: is the zero check really working?
			if query.StartTime == 0 {
				query.StartTime = req.StartTime
			}
			if query.EndTime == 0 {
				query.EndTime = query.EndTime
			}
			// merge it
			// http://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
			results = append(results, readSvc.QueryInt(query)...)
		}

		return res, nil
	}
}

func (ReadServiceHTTPFactory) MakeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var req readRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}
		return req, nil
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
