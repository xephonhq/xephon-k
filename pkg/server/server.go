package server

import (
	"context"
	"net/http"

	"encoding/json"

	httptransport "github.com/go-kit/kit/transport/http"
)

type Server struct {
}

func (Server) Start() {
	// FIXME: the context part in html is also outdated
	// ctx := context.Background()
	infoSvc := infoService{}

	infoHandler := httptransport.NewServer(
		// ctx,
		makeInfoEndpoint(infoSvc),
		decodeInfoRequest,
		encodeResponse,
	)

	http.Handle("/info", infoHandler)
	log.Infof("start serving on 0.0.0.0:%d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return infoRequest{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
