package http

import (
	"net/http"

	"encoding/json"
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	pb "github.com/xephonhq/xephon-k/pkg/server/payload"
	"time"
)

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
