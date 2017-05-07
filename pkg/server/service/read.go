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
	QueryInt(queries []common.Query) ([]common.QueryResult, []common.IntSeries, error)
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
	Error        bool                 `json:"error"`
	ErrorMsg     string               `json:"error_msg"`
	QueryResults []common.QueryResult `json:"query_results"`
	// TODO: where is the data?
	Metrics []common.IntSeries `json:"metrics"`
}

type ReadServiceHTTPFactory struct {
}

func (ReadServiceHTTPFactory) MakeEndpoint(service Service) endpoint.Endpoint {
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
		// apply all the top level query criteria to all the other queries
		// and check if each query is valid
		for i := 0; i < len(req.Queries); i++ {
			// we need to modify it, so we MUST use a pointer, otherwise we are operating on the copy
			// see playground/range_test.go TestRange_Modify
			query := &req.Queries[i]
			if query.StartTime == 0 {
				if req.StartTime == 0 {
					return res, errors.Errorf("%d query lacks start time", i)
				}
				query.StartTime = req.StartTime
			}
			if query.EndTime == 0 {
				if req.EndTime == 0 {
					return res, errors.Errorf("%d query lacks end time", i)
				}
				query.EndTime = req.EndTime
			}
			// aggregator is not required, but we will apply the global aggregator if single query does not have one
			if query.Aggregator.Type == "" && req.Aggregator.Type != "" {
				query.Aggregator = req.Aggregator
			}
			// TODO: handle __name__, and also add __name__ to tag would result in different hash of SeriesID
		}

		queryResults, series, err := readSvc.QueryInt(req.Queries)
		res.QueryResults = queryResults
		res.Metrics = series

		return res, err
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
func (rs ReadServiceServerImpl) QueryInt(queries []common.Query) ([]common.QueryResult, []common.IntSeries, error) {
	return rs.store.QueryIntSeriesBatch(queries)
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
