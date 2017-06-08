package grpc

import (
	"net"

	"fmt"
	"github.com/pkg/errors"
	pb "github.com/xephonhq/xephon-k/pkg/server/payload"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"github.com/xephonhq/xephon-k/pkg/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log = util.Logger.NewEntryWithPkg("k.server.grpc")

type Server struct {
	g      *grpc.Server
	config Config
	write  *service.WriteService2
}

func NewServer(config Config, write *service.WriteService2) *Server {
	return &Server{
		config: config,
		write:  write,
	}
}

// FIXME: the new context in go 1.7?
// https://github.com/grpc/grpc-go/issues/711
func (s *Server) Write(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	// TODO: is there any information I can log from its context?
	return s.write.Write(req)
}

func (s *Server) Start() error {
	// TODO: shutdown the server?
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	t, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrapf(err, "can't start tcp server for grpc on %s", addr)
	}
	s.g = grpc.NewServer()
	pb.RegisterWriteServer(s.g, s)
	reflection.Register(s.g)
	log.Infof("serve grpc on %s", addr)
	if err := s.g.Serve(t); err != nil {
		return errors.Wrapf(err, "can't start grpc server on %s", addr)
	}
	return nil
}

func (s *Server) Stop() {
	// TODO: add a timeout
	log.Info("stopping grpc server")
	// FIXME: #56 this is nil because the net.Listen has not finished, when other error, like http server failed already happened
	// a problem is the tcp server and the grpc server may still try to create go routine leak? ... we can have a mutex
	if s.g == nil {
		log.Warn("grpc server is not even started, but asked to stop")
		return
	}
	s.g.GracefulStop()
	log.Info("grpc server stopped")
}
