package server

import (
	"net/http/httptest"
	"testing"
)

func TestHTTPServer_Mux(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	srv := HTTPServer{Backend: "memory"}
	ts := httptest.NewServer(srv.Mux())
	defer ts.Close()
}
