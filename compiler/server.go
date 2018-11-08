package compiler

import (
	"context"
	"fmt"
	"net"

	svc "github.com/nlepage/codyglot/service/compiler"
	fssvc "github.com/nlepage/codyglot/service/filestore"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Server is a struct implementing CompilerServer as a specific compiler
type Server struct {
	fn     func(context.Context, *fssvc.Id) (*svc.CompileResult, error)
	config ServerConfig
}

// Compile calls Server.fn
func (s *Server) Compile(ctx context.Context, fsID *fssvc.Id) (*svc.CompileResult, error) {
	return s.fn(ctx, fsID)
}

var _ svc.CompilerServer = (*Server)(nil)

// Serve starts listening for gRPC requests
func (s *Server) Serve() error {
	log.Infoln("Starting compiler on port", s.config.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return errors.Wrap(err, "Server: Failed to listen")
	}

	grpcSrv := grpc.NewServer()
	svc.RegisterCompilerServer(grpcSrv, s)
	if err := grpcSrv.Serve(lis); err != nil {
		return errors.Wrap(err, "Server: Failed to serve")
	}

	return nil
}
