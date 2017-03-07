package server

// NOTE: the example on go kit website is outdated
import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/xephonhq/xephon-k/pkg"
)

type InfoService interface {
	Version() string
}

type infoService struct {
}

func (infoService) Version() string {
	return pkg.Version
}

type infoRequest struct {
}

type infoResponse struct {
	Version string `json:"version"`
}

func makeInfoEndpoint(svc InfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// TODO: why there is a cast like this
		// _ := request.(infoRequest)
		v := svc.Version()
		return infoResponse{Version: v}, nil
	}
}
