package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"
	"github.com/xephonhq/xephon-k/pkg/storage/disk"
)

type WriteService interface {
	Service
	WriteInt([]common.IntSeries) error
	WriteDouble([]common.DoubleSeries) error
}

var _ WriteService = (*WriteServiceServerImpl)(nil)

// WriteServiceServerImpl is the server implementation of WriteService
type WriteServiceServerImpl struct {
	store storage.Store
}

type writeRequest struct {
	intSeries    []common.IntSeries
	doubleSeries []common.DoubleSeries
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

		if len(req.intSeries) > 0 {
			err := writeSvc.WriteInt(req.intSeries)
			if err != nil {
				res.Error = true
				res.ErrorMsg = err.Error()
				// TODO: maybe we should return error and let upper encode it into json
				return res, nil
			}
		}
		if len(req.doubleSeries) > 0 {
			err := writeSvc.WriteDouble(req.doubleSeries)
			if err != nil {
				res.Error = true
				res.ErrorMsg = err.Error()
				// TODO: maybe we should return error and let upper encode it into json
				return res, nil
			}
		}
		return res, nil
	}
}

func (WriteServiceHTTPFactory) MakeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var rawSeries []common.RawSeries
		var intSeries []common.IntSeries
		var doubleSeries []common.DoubleSeries
		// FIXME: go-kit does not handle decode error?
		// https://github.com/xephonhq/xephon-k/issues/6
		// https://github.com/go-kit/kit/issues/133
		if err := json.NewDecoder(r.Body).Decode(&rawSeries); err != nil {
			return nil, errors.Wrap(err, "can't decode write request into meta series")
		}
		totalSeries := len(rawSeries)
		log.Tracef("got %d meta series after decode ", len(rawSeries))
		for i := 0; i < totalSeries; i++ {
			switch rawSeries[i].GetSeriesType() {
			case common.TypeIntSeries:
				// copy the meta and decode the points
				s := common.IntSeries{
					SeriesMeta: rawSeries[i].GetMetaCopy(),
				}
				points := make([]common.IntPoint, 0)
				err := json.Unmarshal(rawSeries[i].Points, &points)
				if err != nil {
					return writeRequest{}, errors.Wrapf(err, "can't decode %s into int series", s.Name)
				}
				s.Points = points
				intSeries = append(intSeries, s)
			case common.TypeDoubleSeries:
				s := common.DoubleSeries{
					SeriesMeta: rawSeries[i].GetMetaCopy(),
				}
				points := make([]common.DoublePoint, 0)
				err := json.Unmarshal(rawSeries[i].Points, &points)
				if err != nil {
					return writeRequest{}, errors.Wrapf(err, "can't decode %s into double series", s.Name)
				}
				s.Points = points
				doubleSeries = append(doubleSeries, s)
			default:
				return writeRequest{}, errors.Errorf("unsupported series type %d", rawSeries[i].GetSeriesType())
			}
		}
		log.Tracef("got %d int series after decode ", len(intSeries))
		log.Tracef("got %d double series after decode ", len(doubleSeries))
		return writeRequest{intSeries: intSeries, doubleSeries: doubleSeries}, nil
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

func NewWriteServiceDisk(dir string) *WriteServiceServerImpl {
	store := disk.GetDefaultDiskStore(dir)
	return &WriteServiceServerImpl{store: store}
}

func NewWriteServiceCassandra(host string) *WriteServiceServerImpl {
	store := cassandra.GetDefaultCassandraStore(host)
	return &WriteServiceServerImpl{store: store}
}

// ServiceName implements Service
func (ws *WriteServiceServerImpl) ServiceName() string {
	return "write"
}

// WriteInt implements WriteService
func (ws *WriteServiceServerImpl) WriteInt(series []common.IntSeries) error {
	// write to memory storage
	// NOTE: maybe we should wrap error instead of just return it
	return ws.store.WriteIntSeries(series)
}

// WriteDouble implements WriteService
func (ws *WriteServiceServerImpl) WriteDouble(series []common.DoubleSeries) error {
	return ws.store.WriteDoubleSeries(series)
}
