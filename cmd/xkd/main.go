package main

import (
	"github.com/xephonhq/xephon-k/pkg/server/http"
)

// The daemon
func main() {
	// temporal workaround for try out new server
	srv := http.NewServer()
	srv.Start()
}
