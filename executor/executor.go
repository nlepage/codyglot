package executor

import (
	"context"
	"fmt"
	"net"

	"github.com/nlepage/codyglot/ping"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Executor is a struct implementing CodyglotServer as a specific executor
type Executor struct {
	config Config
	*ping.Server
	fn        func(context.Context, *service.ExecuteRequest) (*service.ExecuteResponse, error)
	languages []string
}

// New creates a new Executor
func New(config Config, fn func(context.Context, *service.ExecuteRequest) (*service.ExecuteResponse, error), languages []string) *Executor {
	return &Executor{config, &ping.Server{}, fn, languages}
}

// Execute calls Executor.fn
func (e *Executor) Execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	return e.fn(ctx, req)
}

// Languages lists languages supported by the executor
func (e *Executor) Languages(ctx context.Context, req *service.LanguagesRequest) (*service.LanguagesResponse, error) {
	return &service.LanguagesResponse{Languages: e.languages}, nil
}

var _ service.CodyglotServer = (*Executor)(nil)

// Serve starts listening for gRPC requests
func (e *Executor) Serve() error {
	log.Infoln("Starting executor on port", e.config.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", e.config.Port))
	if err != nil {
		return errors.Wrap(err, "Executor: Failed to listen")
	}

	grpcSrv := grpc.NewServer()
	service.RegisterCodyglotServer(grpcSrv, e)
	if err := grpcSrv.Serve(lis); err != nil {
		return errors.Wrap(err, "Executor: Failed to serve")
	}

	return nil
}
