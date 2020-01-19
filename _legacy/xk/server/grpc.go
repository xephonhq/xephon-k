package server

import (
	"context"

	dlog "github.com/dyweb/gommon/log"

	pb "github.com/xephonhq/xephon-k/xk/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb.XephonkServer = (*GrpcServer)(nil)

type GrpcServer struct {
	log *dlog.Logger
}

func NewGrpcServer() (*GrpcServer, error) {
	srv := &GrpcServer{}
	dlog.NewStructLogger(log, srv)
	return srv, nil
}

func (srv *GrpcServer) Ping(ctx context.Context, ping *pb.PingReq) (*pb.PingRes, error) {
	return &pb.PingRes{Message: "pong from xephonk your message is " + ping.Message}, nil
}

func (srv *GrpcServer) WritePoints(ctx context.Context, req *pb.WritePointsReq) (*pb.WritePointsRes, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (srv *GrpcServer) WriteSeries(ctx context.Context, req *pb.WriteSeriesReq) (*pb.WriteSeriesRes, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (srv *GrpcServer) PrepareSeries(ctx context.Context, req *pb.PrepareSeriesReq) (*pb.PrepareSeriesRes, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (srv *GrpcServer) WritePreparedPoints(ctx context.Context, req *pb.WritePreparedPointsReq) (*pb.WritePreparedPointsRes, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (srv *GrpcServer) WritePreparedSeries(ctx context.Context, req *pb.WritePreparedSeriesReq) (*pb.WritePreparedSeriesRes, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (srv *GrpcServer) WritePreparedSeriesColumnar(ctx context.Context, req *pb.WritePreparedSeriesColumnarReq) (*pb.WritePreparedSeriesColumnarRes, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
