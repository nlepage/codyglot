package server

import (
	"context"
	"fmt"
	"net"

	executor "github.com/nlepage/codyglot/executor/service"
	"github.com/nlepage/codyglot/router/server/config"
	router "github.com/nlepage/codyglot/router/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Server is Codyglot router gRPC server
type Server struct{}

// Init initializes the router server
func Init() error {
	return config.InitLanguagesMap()
}

// Serve starts listening for gRPC requests
func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return fmt.Errorf("Failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	router.RegisterRouterServer(grpcSrv, s)
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v", err)
	}

	return nil
}

// Execute forwards a code request execution to an executor
func (s *Server) Execute(ctx context.Context, req *router.ExecuteRequest) (*router.ExecuteResponse, error) {
	endpoint, ok := config.LanguagesMap[req.Language]
	if !ok {
		return nil, fmt.Errorf("Config: No executor endpoint for language %s", req.Language)
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Router: could not create executor dial context")
	}
	defer conn.Close()

	client := executor.NewExecutorClient(conn)

	res, err := client.Execute(ctx, &executor.ExecuteRequest{
		Source: req.Source,
		Stdin:  req.Stdin,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Router: an error occured while executing request to executor")
	}

	return &router.ExecuteResponse{
		ExitStatus:      res.ExitStatus,
		Stderr:          res.Stderr,
		Stdout:          res.Stdout,
		CompilationTime: res.CompilationTime,
		RunningTime:     res.RunningTime,
	}, nil
}

// Languages lists languages for which an executor is available
func (s *Server) Languages(ctx context.Context, req *router.LanguagesRequest) (*router.LanguagesResponse, error) {
	res := router.LanguagesResponse{
		Languages: make([]string, 0, len(config.LanguagesMap)),
	}

	for language := range config.LanguagesMap {
		res.Languages = append(res.Languages, language)
	}

	return &res, nil
}

var _ router.RouterServer = (*Server)(nil)
