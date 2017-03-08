package server

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg/server/middleware"
	"github.com/xephonhq/xephon-k/pkg/server/service"
)

type Server struct {
}

func (Server) Start() {
	var infoSvc service.InfoService
	infoSvc = service.InfoServiceImpl{}
	infoSvc = middleware.NewLoggingInfoServiceMiddleware(infoSvc)
	infoSvcHTTPFactory := service.InfoServiceHTTPFactory{}

	infoHandler := httptransport.NewServer(
		infoSvcHTTPFactory.MakeEndpoint(infoSvc),
		infoSvcHTTPFactory.MakeDecode(),
		infoSvcHTTPFactory.MakeEncode(),
	)

	var writeSvc service.WriteService
	writeSvc = service.WriteServiceImpl{}
	writeSvc = middleware.NewLoggingWriteServiceMiddleware(writeSvc)
	writeSvcHTTPFactory := service.WriteServiceHTTPFactory{}

	writeHandler := httptransport.NewServer(
		writeSvcHTTPFactory.MakeEndpoint(writeSvc),
		writeSvcHTTPFactory.MakeDecode(),
		writeSvcHTTPFactory.MakeEncode(),
	)

	http.Handle("/info", infoHandler)
	http.Handle("/write", writeHandler)
	log.Infof("start serving on 0.0.0.0:%d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
