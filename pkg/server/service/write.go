package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg/common"
)

type WriteService interface {
	Service
	WriteInt([]common.IntSeries) error
}

// WriteServiceImpl is the server implementation of WriteService
type WriteServiceImpl struct {
}

type writeRequest struct {
	series []common.IntSeries
}

type writeResponse struct {
}

type WriteServiceHTTPFactory struct {
}

func (WriteServiceHTTPFactory) MakeEndpoint(service Service) endpoint.Endpoint {
	writeScv, ok := service.(WriteService)
	if !ok {
		log.Panic("must pass write service to write service factory")
	}
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(writeRequest)
		if !ok {
			log.Panic("should be writeRequest")
		}
		err := writeScv.WriteInt(req.series)
		if err != nil {
			return nil, err
		}
		return writeResponse{}, nil
	}
}

func (WriteServiceHTTPFactory) MakeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		// TODO: do I need to allocate space?
		var series []common.IntSeries
		if err := json.NewDecoder(r.Body).Decode(&series); err != nil {
			return nil, err
		}
		return writeRequest{series: series}, nil
	}
}

func (WriteServiceHTTPFactory) MakeEncode() httptransport.EncodeResponseFunc {
	return httptransport.EncodeJSONResponse
}

func (WriteServiceImpl) ServiceName() string {
	return "write"
}

func (ws WriteServiceImpl) WriteInt(series []common.IntSeries) error {
	return nil
}
