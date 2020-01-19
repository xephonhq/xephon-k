package http

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	"net/http"
	"time"
)

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
	Metrics      []common.Series      `json:"metrics"`
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

func (s *Server) read(w http.ResponseWriter, r *http.Request) {
	defer func(begin time.Time) {
		log.Infof("POST /read %d", time.Since(begin))
	}(time.Now())
	var req readRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		invalidFormat(w, errors.Wrap(err, "can't parse read request"))
		return
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
				invalidFormat(w, errors.Errorf("%d query lacks start time", i))
				return
			}
			query.StartTime = req.StartTime
		}
		if query.EndTime == 0 {
			if req.EndTime == 0 {
				invalidFormat(w, errors.Errorf("%d query lacks end time", i))
				return
			}
			query.EndTime = req.EndTime
		}
		// aggregator is not required, but we will apply the global aggregator if single query does not have one
		if query.Aggregator.Type == "" && req.Aggregator.Type != "" {
			query.Aggregator = req.Aggregator
		}
		// TODO: handle __name__, and also add __name__ to tag would result in different hash of SeriesID
	}

	queryResults, series, err := s.readSvc.QuerySeries(req.Queries)
	if err != nil {
		internalError(w, err)
		return
	}
	res.QueryResults = queryResults
	res.Metrics = series

	writeJSON(w, res)
	return
}
