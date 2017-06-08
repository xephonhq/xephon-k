package service

import (
	pb "github.com/xephonhq/xephon-k/pkg/server/payload"
	"github.com/xephonhq/xephon-k/pkg/storage"
)

type WriteService struct {
	store storage.Store
}

func NewWriteService(store storage.Store) *WriteService {
	return &WriteService{
		store: store,
	}
}

// TODO: support context, but protobuf is using x/net/context
func (w *WriteService) Write(req *pb.WriteRequest) (*pb.WriteResponse, error) {
	res := &pb.WriteResponse{Error: false, ErrorMsg: ""}

	if len(req.IntSeries) > 0 {
		err := w.store.WriteIntSeries(req.IntSeries)
		if err != nil {
			res.Error = true
			res.ErrorMsg = err.Error()
			// TODO: maybe we should return error and let upper encode it into json
			return res, nil
		}
	}
	if len(req.DoubleSeries) > 0 {
		err := w.store.WriteDoubleSeries(req.DoubleSeries)
		if err != nil {
			res.Error = true
			res.ErrorMsg = err.Error()
			// TODO: maybe we should return error and let upper encode it into json
			return res, nil
		}
	}
	return res, nil
}
