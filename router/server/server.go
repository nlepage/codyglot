package server

import (
	"context"
	"fmt"
	"net"

	executor "github.com/Zenika/codyglot/executor/service"
	router "github.com/Zenika/codyglot/router/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Server is Codyglot router gRPC server
type Server struct {
	LanguagesMap map[string]string
	Port         int
}

// Serve starts listening for gRPC requests
func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
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
	target, ok := s.LanguagesMap[req.Language]
	if !ok {
		return nil, fmt.Errorf("Router: No executor for language %s", req.Language)
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Router: could not create executor dial context")
	}
	defer conn.Close()

	client := executor.NewExecutorClient(conn)

	res, err := client.Execute(ctx, &executor.ExecuteRequest{
		Code:  req.Code,
		StdIn: req.StdIn,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Router: an error occured while executing request to executor")
	}

	return &router.ExecuteResponse{
		ExitCode: res.ExitCode,
		StdErr:   res.StdErr,
		StdOut:   res.StdOut,
	}, nil
}

// Languages lists languages for which an executor is available
func (s *Server) Languages(ctx context.Context, req *router.LanguagesRequest) (*router.LanguagesResponse, error) {
	res := router.LanguagesResponse{
		Languages: make([]string, 0, len(s.LanguagesMap)),
	}

	for language := range s.LanguagesMap {
		res.Languages = append(res.Languages, language)
	}

	return &res, nil
}

var _ router.RouterServer = (*Server)(nil)
