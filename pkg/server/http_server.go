package server

import (
	"net/http"

	"context"
	"encoding/json"
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

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// TODO: set the status code based on type of error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":     true,
		"error_msg": err.Error(),
	})
}

func (srv HTTPServer) Mux() *http.ServeMux {
	options := []httptransport.ServerOption{
		// TODO: replace logger as well, but go-kit's log interface is quite strange
		// it suggests structured logging, which is key-value for everything, and by default just one level
		// the default logger is a nop logger
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	var infoSvc service.InfoService
	infoSvc = service.InfoServiceImpl{}
	infoSvc = middleware.NewLoggingInfoServiceMiddleware(infoSvc)
	infoSvcHTTPFactory := service.InfoServiceHTTPFactory{}

	infoHandler := httptransport.NewServer(
		infoSvcHTTPFactory.MakeEndpoint(infoSvc),
		infoSvcHTTPFactory.MakeDecode(),
		infoSvcHTTPFactory.MakeEncode(),
		options...,
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
		options...,
	)

	readSvc = middleware.NewLoggingReadServiceMiddleware(readSvc)
	readSvcHTTPFactory := service.ReadServiceHTTPFactory{}

	readHandler := httptransport.NewServer(
		readSvcHTTPFactory.MakeEndpoint(readSvc),
		readSvcHTTPFactory.MakeDecode(),
		readSvcHTTPFactory.MakeEncode(),
		options...,
	)

	mux := http.NewServeMux()
	mux.Handle("/info", infoHandler)
	mux.Handle("/write", writeHandler)
	mux.Handle("/w", writeHandler)
	mux.Handle("/read", readHandler)
	mux.Handle("/r", readHandler)
	// TODO: allow config static content location
	// TODO: accessing static content is not logged
	mux.Handle("/", http.FileServer(http.Dir(".")))
	return mux
}

func (srv HTTPServer) Start() {
	log.Infof("start serving on 0.0.0.0:%d", srv.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", srv.Port), srv.Mux()))
}
