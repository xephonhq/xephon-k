package grpc

import (
	"net"

	pb "github.com/xephonhq/xephon-k/pkg/server/payload"
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log = util.Logger.NewEntryWithPkg("k.server.grpc")

type Server struct {
	store storage.Store
}

// FIXME: the new context in go 1.7?
// https://github.com/grpc/grpc-go/issues/711
func (s *Server) Write(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	return &pb.WriteResponse{Error: false, ErrorMsg: ""}, nil
}

func (s *Server) Start() {
	// TODO: config, kip the reference and shutdown the server?
	t, err := net.Listen("tcp", "localhost:2444")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	pb.RegisterWriteServer(gs, s)
	reflection.Register(gs)
	if err := gs.Serve(t); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
