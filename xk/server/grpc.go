package server

import (
	dlog "github.com/dyweb/gommon/log"

	"context"
	rpc "github.com/xephonhq/xephon-k/xk/transport/grpc"
	pb "github.com/xephonhq/xephon-k/xk/xkpb"
)

var _ rpc.XephonkServer = (*GrpcServer)(nil)

type GrpcServer struct {
	log *dlog.Logger
}

func NewGrpcServer() (*GrpcServer, error) {
	srv := &GrpcServer{}
	dlog.NewStructLogger(log, srv)
	return srv, nil
}

func (srv *GrpcServer) Ping(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	return &pb.Pong{Message: "pong from xephonk your message is " + ping.Message}, nil
}
