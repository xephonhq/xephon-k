package server

import (
	"net/http"

	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/xephonhq/xephon-k/pkg/server/middleware"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"strings"
)

type HTTPServer struct {
	Port          int
	Backend       string
	CassandraHost string
}

func (srv HTTPServer) Mux() *http.ServeMux {
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
		writeSvc = service.NewWriteServiceCassandra(srv.CassandraHost)
		readSvc = service.NewReadServiceCassandra(srv.CassandraHost)
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

	mux := http.NewServeMux()
	mux.Handle("/info", infoHandler)
	mux.Handle("/write", writeHandler)
	mux.Handle("/w", writeHandler)
	mux.Handle("/read", readHandler)
	mux.Handle("/r", readHandler)
	return mux
}

func (srv HTTPServer) Start() {
	log.Infof("start serving on 0.0.0.0:%d", srv.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", srv.Port), srv.Mux()))
}
