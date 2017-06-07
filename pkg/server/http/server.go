package http

import (
	"fmt"
	"net/http"

	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.server.http")

type Server struct {
	store storage.Store
}

// TODO: functional style config and config storage
func NewServer() *Server {
	return &Server{}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)

	log.Info("start serving on localhost:2333")
	http.ListenAndServe("localhost:2333", mux)
}
