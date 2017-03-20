package server

import (
	"net/http"

	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg/server/middleware"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"strings"
)

type Server struct {
	Port    int
	Backend string
}

func (srv Server) Start() {
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
	var readSvc service.ReadService

	if strings.HasPrefix(srv.Backend, "m") {
		log.Info("use memory backend")
		writeSvc = service.NewWriteServiceMem()
		readSvc = service.NewReadServiceMem()
	} else if strings.HasPrefix(srv.Backend, "c") {
		log.Info("use cassandra backend")
		writeSvc = service.NewWriteServiceCassandra()
		readSvc = service.NewReadServiceCassandra()
	} else {
		log.Fatalf("unknown backend %s", srv.Backend)
	}

	writeSvc = middleware.NewLoggingWriteServiceMiddleware(writeSvc)
	writeSvcHTTPFactory := service.WriteServiceHTTPFactory{}

	writeHandler := httptransport.NewServer(
		writeSvcHTTPFactory.MakeEndpoint(writeSvc),
		writeSvcHTTPFactory.MakeDecode(),
		writeSvcHTTPFactory.MakeEncode(),
	)

	readSvc = middleware.NewLoggingReadServiceMiddleware(readSvc)
	readSvcHTTPFactory := service.ReadServiceHTTPFactory{}

	readHandler := httptransport.NewServer(
		readSvcHTTPFactory.MakeEndpoint(readSvc),
		readSvcHTTPFactory.MakeDecode(),
		readSvcHTTPFactory.MakeEncode(),
	)

	http.Handle("/info", infoHandler)
	http.Handle("/write", writeHandler)
	http.Handle("/read", readHandler)
	log.Infof("start serving on 0.0.0.0:%d", srv.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", srv.Port), nil))
}
