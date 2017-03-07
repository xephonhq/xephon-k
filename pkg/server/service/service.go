package service

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Service is used as a workaround for letting makeEndpoint work for all type of services
type Service interface {
	ServiceName() string
}

// Factory is the base factory
type Factory interface {
	MakeEndpoint(service Service) endpoint.Endpoint
}

// HTTPFactory is used when building a HTTP Server
type HTTPFactory interface {
	Factory
	MakeDecode() httptransport.DecodeRequestFunc
	MakeEncode() httptransport.EncodeResponseFunc
}
