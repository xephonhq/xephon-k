package server

// NOTE: the example on go kit website is outdated
import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg"
)

type InfoService interface {
	Service
	Version() string
}

type infoRequest struct {
}

type infoResponse struct {
	Version string `json:"version"`
}

type infoServiceHTTPFactory struct {
}

func (infoServiceHTTPFactory) makeEndpoint(service Service) endpoint.Endpoint {
	infoSvc, ok := service.(infoService)
	if !ok {
		log.Panic("must pass info service to info service factory")
	}
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		v := infoSvc.Version()

		return infoResponse{Version: v}, nil
	}
}

func (infoServiceHTTPFactory) makeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		return infoRequest{}, nil
	}
}

func (infoServiceHTTPFactory) makeEncode() httptransport.EncodeResponseFunc {
	return encodeResponse
}

type infoService struct {
}

func (infoService) ServiceName() string {
	return "info"
}

func (infoService) Version() string {
	return pkg.Version
}

func makeInfoEndpoint(svc InfoService) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		v := svc.Version()
		return infoResponse{Version: v}, nil
	}
}

func decodeInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return infoRequest{}, nil
}
