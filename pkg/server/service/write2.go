package service

import (
	pb "github.com/xephonhq/xephon-k/pkg/server/payload"
	"github.com/xephonhq/xephon-k/pkg/storage"
)

type WriteService2 struct {
	store storage.Store
}

func NewWriteService(store storage.Store) *WriteService2 {
	return &WriteService2{
		store: store,
	}
}

// TODO: support context, but protobuf is using x/net/context
func (w *WriteService2) Write(req *pb.WriteRequest) (*pb.WriteResponse, error) {
	return &pb.WriteResponse{Error: false, ErrorMsg: ""}, nil
}
