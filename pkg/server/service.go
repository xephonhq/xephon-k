package server

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Service is used as a workaround for letting makeEndpoint work for all type of services
type Service interface {
	ServiceName() string
}

// ServiceFactory is the base factory
type ServiceFactory interface {
	makeEndpoint(service Service) endpoint.Endpoint
}

// type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

// type decodeHTTPRequest func(ctx context.Context, r *http.Request) (interface{}, error)
// type encodeHTTPResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error

// ServiceHTTPFactory is used when building a HTTP Server
type ServiceHTTPFactory interface {
	ServiceFactory
	makeDecode() httptransport.DecodeRequestFunc
	makeEncode() httptransport.EncodeResponseFunc
}
