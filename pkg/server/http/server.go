package http

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.server.http")

type Server struct {
	h        *http.Server
	config   Config
	writeSvc *service.WriteService
	readSvc  *service.ReadService
}

func NewServer(config Config, write *service.WriteService, read *service.ReadService) *Server {
	return &Server{
		config:   config,
		writeSvc: write,
		readSvc:  read,
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func info(w http.ResponseWriter, r *http.Request) {
	// TODO: information from the store
	writeJSON(w, map[string]string{"version": pkg.Version})
}

func writeErr(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":     true,
		"error_msg": err.Error(),
	})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// TODO: check error?
	json.NewEncoder(w).Encode(v)
}

func invalidFormat(w http.ResponseWriter, err error) {
	writeErr(w, err, http.StatusBadRequest)
}

func internalError(w http.ResponseWriter, err error) {
	writeErr(w, err, http.StatusInternalServerError)
}

func (s *Server) Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/info", info)
	mux.HandleFunc("/write", s.write)
	mux.HandleFunc("/read", s.read)

	if s.config.EnablePProf {
		log.Info("pprof is enabled")
		// TODO: it seems it prevent the http server from being stopped gracefully
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	return mux
}

func (s *Server) Start() error {
	mux := s.Mux()

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	s.h = &http.Server{Addr: addr, Handler: mux}
	log.Infof("serve http on %s", addr)
	if err := s.h.ListenAndServe(); err != nil {
		return errors.Wrapf(err, "can't start http server on %s", addr)
	}
	return nil
}

// Stop tries to shutdown the HTTP server gracefully
// TODO: it seems after pprof is enabled, it can't stop gracefully
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
