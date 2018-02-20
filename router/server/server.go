package server

import (
	"context"
	"fmt"
	"net"

	"github.com/nlepage/codyglot/ping"
	"github.com/nlepage/codyglot/router/server/config"
	"github.com/nlepage/codyglot/router/server/executor"
	"github.com/nlepage/codyglot/service"
	"google.golang.org/grpc"
)

// Init initializes the router server (parses CLI args and starts background tasks)
func Init() {
	executor.Init()
}

// Server is Codyglot router gRPC server
type Server struct {
	*ping.Server
}

// New returns a pointer on a new Server
func New() *Server {
	return &Server{&ping.Server{}}
}

// Serve starts listening for gRPC requests
func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return fmt.Errorf("Failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	service.RegisterCodyglotServer(grpcSrv, s)
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v", err)
	}

	return nil
}

// Execute forwards a code request execution to an executor
func (s *Server) Execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	return executor.Execute(ctx, req)
}

// Languages lists languages for which an executor is available
func (s *Server) Languages(ctx context.Context, req *service.LanguagesRequest) (*service.LanguagesResponse, error) {
	return &service.LanguagesResponse{
		Languages: executor.Languages(),
	}, nil
}

var _ service.CodyglotServer = (*Server)(nil)
