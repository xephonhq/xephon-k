package service

// NOTE: the example on go kit website is outdated
import (
	"context"
	"log"
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

type InfoServiceHTTPFactory struct {
}

func (InfoServiceHTTPFactory) MakeEndpoint(service Service) endpoint.Endpoint {
	infoSvc, ok := service.(InfoService)
	if !ok {
		log.Panic("must pass info service to info service factory")
	}
	// FIXME: the naming here is misleading, the info actually return all the info, more than just version
	// and how to hand things like info/version in go-kit
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		v := infoSvc.Version()

		return infoResponse{Version: v}, nil
	}
}

func (InfoServiceHTTPFactory) MakeDecode() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		return infoRequest{}, nil
	}
}

func (InfoServiceHTTPFactory) MakeEncode() httptransport.EncodeResponseFunc {
	return httptransport.EncodeJSONResponse
}

type InfoServiceImpl struct {
}

func (InfoServiceImpl) ServiceName() string {
	return "info"
}

func (InfoServiceImpl) Version() string {
	return pkg.Version
}
