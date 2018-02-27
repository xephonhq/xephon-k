package server

import (
	dlog "github.com/dyweb/gommon/log"
	"net/http"
)

type HttpServer struct {
	log *dlog.Logger
}

func NewHttpServer() (*HttpServer, error) {
	s := &HttpServer{}
	dlog.NewStructLogger(log, s)
	return s, nil
}

func (srv *HttpServer) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	return mux
}
