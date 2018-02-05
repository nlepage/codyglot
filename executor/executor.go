package executor

import (
	"context"
	"fmt"
	"net"

	config "github.com/nlepage/codyglot/executor/config"
	"github.com/nlepage/codyglot/executor/tmputil"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Init initializes the executor
func Init() {
	tmputil.StartCleanup()
}

// Executor is a struct implementing CodyglotServer as a specific executor
type Executor struct {
	fn        func(context.Context, *service.ExecuteRequest) (*service.ExecuteResponse, error)
	languages []string
}

// New creates a new Executor
func New(fn func(context.Context, *service.ExecuteRequest) (*service.ExecuteResponse, error), languages []string) *Executor {
	return &Executor{fn, languages}
}

// Execute calls Executor.fn
func (e *Executor) Execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	return e.fn(ctx, req)
}

// Languages lists languages supported by the executor
func (e *Executor) Languages(ctx context.Context, req *service.LanguagesRequest) (*service.LanguagesResponse, error) {
	return &service.LanguagesResponse{e.languages}, nil
}

var _ service.CodyglotServer = (*Executor)(nil)

// Serve starts listening for gRPC requests
func (e *Executor) Serve() error {
	log.Infoln("Starting executor on port", config.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
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
