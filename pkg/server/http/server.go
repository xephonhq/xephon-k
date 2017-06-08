package http

import (
	"fmt"
	"net/http"

	"context"
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"github.com/xephonhq/xephon-k/pkg/util"
	"time"
)

var log = util.Logger.NewEntryWithPkg("k.server.http")

type Server struct {
	h      *http.Server
	config Config
	write  *service.WriteService2
}

// TODO: functional style config and config storage
func NewServer(config Config, write *service.WriteService2) *Server {
	return &Server{
		config: config,
		write:  write,
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	s.h = &http.Server{Addr: addr, Handler: mux}
	log.Infof("serve http on %s", addr)
	if err := s.h.ListenAndServe(); err != nil {
		return errors.Wrapf(err, "can't start http server on %s", addr)
	}
	return nil
}

// TODO: graceful shutdown, need to store server
// https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2
func (s *Server) Stop() {
	log.Info("stopping http server")
	if s.h == nil {
		log.Warn("http server is not even started, but asked to stop")
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := s.h.Shutdown(ctx); err != nil {
		log.Warnf("didn't stop gracefully in 5s %v", err)
	}
	log.Info("http server stopped")
}
