package server

import (
	"context"
	"sync"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	ihttp "github.com/at15/go.ice/ice/transport/http"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc"

	"github.com/xephonhq/xephon-k/xk/config"
	mygrpc "github.com/xephonhq/xephon-k/xk/transport/grpc"
)

type Manager struct {
	cfg config.ServerConfig

	grpcSrv       *GrpcServer
	grpcTransport *igrpc.Server
	httpSrv       *HttpServer
	httpTransport *ihttp.Server

	log *dlog.Logger
}

func NewManager(cfg config.ServerConfig) (*Manager, error) {
	grpcSrv, err := NewGrpcServer()
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc server")
	}
	grpcTransport, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
		mygrpc.RegisterXephonkServer(s, grpcSrv)
	})
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc transport")
	}
	httpSrv, err := NewHttpServer()
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create http server")
	}
	httpTransport, err := ihttp.NewServer(cfg.Http, httpSrv.Handler(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create http transport")
	}
	mgr := &Manager{
		cfg:           cfg,
		grpcSrv:       grpcSrv,
		grpcTransport: grpcTransport,
		httpSrv:       httpSrv,
		httpTransport: httpTransport,
	}
	dlog.NewStructLogger(log, mgr)
	return mgr, nil
}

func (mgr *Manager) Run() error {
	var (
		wg      sync.WaitGroup
		grpcErr error
		httpErr error
		merr    = errors.NewMultiErrSafe()
	)
	wg.Add(2) // grpc + http
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// grpc server
	go func() {
		go func() {
			if err := mgr.grpcTransport.Run(); err != nil {
				grpcErr = err
				cancel()
			}
		}()
		select {
		case <-ctx.Done():
			if grpcErr != nil {
				merr.Append(grpcErr)
				mgr.log.Errorf("can't run grpc server %v", grpcErr)
			} else {
				mgr.log.Warn("TODO: other's fault, need to shutdown grpc server")
			}
			wg.Done()
			return
		}
	}()
	// http server
	go func() {
		go func() {
			if err := mgr.httpTransport.Run(); err != nil {
				httpErr = err
				cancel()
			}
		}()
		select {
		case <-ctx.Done():
			if httpErr != nil {
				merr.Append(httpErr)
				mgr.log.Errorf("can't run http server %v", httpErr)
			} else {
				// other service's fault
				mgr.log.Warn("TODO: other's fault, need to shutdown http server")
			}
			wg.Done()
			return
		}
	}()
	wg.Wait()
	return merr.ErrorOrNil()
}
