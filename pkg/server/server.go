package server

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg/server/service"
)

type Server struct {
}

func (Server) Start() {
	infoSvc := service.InfoServiceImpl{}
	infoSvcHTTPFactory := service.InfoServiceHTTPFactory{}

	infoHandler := httptransport.NewServer(
		infoSvcHTTPFactory.MakeEndpoint(infoSvc),
		infoSvcHTTPFactory.MakeDecode(),
		infoSvcHTTPFactory.MakeEncode(),
	)

	http.Handle("/info", infoHandler)
	log.Infof("start serving on 0.0.0.0:%d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
