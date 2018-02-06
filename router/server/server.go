package server

import (
	"context"
	"fmt"
	"net"

	"github.com/nlepage/codyglot/ping"
	"github.com/nlepage/codyglot/router/server/config"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func Init() error {
	if err := initExecutorsStatic(); err != nil {
		return err
	}

	startPinging()

	return nil
}

// Server is Codyglot router gRPC server
type Server struct {
	*ping.Server
}

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
	endpoint := "" // FIXME

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Router: could not create executor dial context")
	}
	defer conn.Close()

	client := service.NewCodyglotClient(conn)

	res, err := client.Execute(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Router: an error occured while calling executor")
	}

	return res, nil
}

// Languages lists languages for which an executor is available
func (s *Server) Languages(ctx context.Context, req *service.LanguagesRequest) (*service.LanguagesResponse, error) {
	return &service.LanguagesResponse{nil}, nil
}

var _ service.CodyglotServer = (*Server)(nil)
