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
	QueriesRaw json.RawMessage   `json:"queries,omitempty"`
	Queries    []common.Query    `json:"queries_that_cant_be_directly_unmsharl_to"`
	StartTime  int64             `json:"start_time,omitempty"`
	EndTime    int64             `json:"end_time,omitempty"`
	Aggregator common.Aggregator `json:"aggregator,omitemoty"`
}

// for avoid recursion in UnmarshalJSON
type readRequestAlias readRequest

type readResponse struct {
	Error    bool                 `json:"error"`
	ErrorMsg string               `json:"error_msg"`
	Queries  []common.QueryResult `json:"query_results"`
	// TODO: where is the data?
	Metrics []common.IntSeries `json:"metrics"`
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
			// TODO: the logic here should be changed, should not call QueryInt one by one, instead, should
			// pass all the queries to it and let it handle the logic
			results = append(results, readSvc.QueryInt(query)...)
		}
		res.Metrics = results
		return res, nil
	}
}

func (ReadServiceHTTPFactory) MakeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var req readRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, errors.Wrap(err, "can't parse read request")
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

// QueryInt implements ReadService
func (rs ReadServiceServerImpl) QueryInt(q common.Query) []common.IntSeries {
	series, err := rs.store.QueryIntSeries(q)
	// TODO: better error handling, i.e propagate the error to upper level
	if err != nil {
		log.Warn(err)
	}
	return series
}

// UnmarshalJSON implements Unmarshaler interface
func (req *readRequest) UnmarshalJSON(data []byte) error {
	// NOTE: need to use Alias to avoid recursion
	// http://choly.ca/post/go-json-marshalling/
	// http://stackoverflow.com/questions/29667379/json-unmarshal-fails-when-embedded-type-has-
	a := (*readRequestAlias)(req)
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}
	err = json.Unmarshal(req.QueriesRaw, &req.Queries)
	if err != nil {
		return errors.Wrap(err, "queries field is not provided")
	}
	return nil
}
