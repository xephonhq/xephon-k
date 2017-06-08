package http

import (
	"fmt"
	"net/http"

	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	pb "github.com/xephonhq/xephon-k/pkg/server/payload"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"github.com/xephonhq/xephon-k/pkg/util"
	"time"
)

var log = util.Logger.NewEntryWithPkg("k.server.http")

type Server struct {
	h        *http.Server
	config   Config
	writeSvc *service.WriteService2
}

// TODO: functional style config and config storage
func NewServer(config Config, write *service.WriteService2) *Server {
	return &Server{
		config:   config,
		writeSvc: write,
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
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
	json.NewEncoder(w).Encode(v)
}

func invalidFormat(w http.ResponseWriter, err error) {
	writeErr(w, err, http.StatusBadRequest)
}

func internalError(w http.ResponseWriter, err error) {
	writeErr(w, err, http.StatusInternalServerError)
}

func (s *Server) write(w http.ResponseWriter, r *http.Request) {
	defer func(begin time.Time) {
		log.Infof("POST /write %d", time.Since(begin))
	}(time.Now())
	// decode
	var rawSeries []common.RawSeries
	var intSeries []common.IntSeries
	var doubleSeries []common.DoubleSeries
	if err := json.NewDecoder(r.Body).Decode(&rawSeries); err != nil {
		invalidFormat(w, errors.Wrap(err, "can't decode write request into meta series"))
		return
	}
	totalSeries := len(rawSeries)
	log.Tracef("got %d meta series after decode ", len(rawSeries))
	for i := 0; i < totalSeries; i++ {
		switch rawSeries[i].GetSeriesType() {
		case common.TypeIntSeries:
			// copy the meta and decode the points
			s := common.IntSeries{
				SeriesMeta: rawSeries[i].GetMetaCopy(),
			}
			points := make([]common.IntPoint, 0)
			err := json.Unmarshal(rawSeries[i].Points, &points)
			if err != nil {
				invalidFormat(w, errors.Wrapf(err, "can't decode %s into int series", s.Name))
				return
			}
			s.Points = points
			intSeries = append(intSeries, s)
		case common.TypeDoubleSeries:
			s := common.DoubleSeries{
				SeriesMeta: rawSeries[i].GetMetaCopy(),
			}
			points := make([]common.DoublePoint, 0)
			err := json.Unmarshal(rawSeries[i].Points, &points)
			if err != nil {
				invalidFormat(w, errors.Wrapf(err, "can't decode %s into double series", s.Name))
				return
			}
			s.Points = points
			doubleSeries = append(doubleSeries, s)
		default:
			invalidFormat(w, errors.Errorf("unsupported series type %d", rawSeries[i].GetSeriesType()))
			return
		}
	}
	// write the series
	log.Tracef("got %d int series after decode ", len(intSeries))
	log.Tracef("got %d double series after decode ", len(doubleSeries))
	req := &pb.WriteRequest{IntSeries: intSeries, DoubleSeries: doubleSeries}
	res, err := s.writeSvc.Write(req)
	if err != nil {
		internalError(w, err)
		return
	}
	writeJSON(w, res)
}

func (s *Server) Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/write", s.write)
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
