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

var (
	languages map[string]string
)

// Init initializes the router server
func Init() error {
	return initLanguages()
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
	endpoint, ok := languages[req.Language]
	if !ok {
		return nil, fmt.Errorf("Config: No executor endpoint for language %s", req.Language)
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Router: could not create executor dial context")
	}
	defer conn.Close()

	client := service.NewCodyglotClient(conn)

	res, err := client.Execute(ctx, &service.ExecuteRequest{
		Source: req.Source,
		Stdin:  req.Stdin,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Router: an error occured while calling executor")
	}

	return &service.ExecuteResponse{
		ExitStatus:      res.ExitStatus,
		Stderr:          res.Stderr,
		Stdout:          res.Stdout,
		CompilationTime: res.CompilationTime,
		RunningTime:     res.RunningTime,
	}, nil
}

// Languages lists languages for which an executor is available
func (s *Server) Languages(ctx context.Context, req *service.LanguagesRequest) (*service.LanguagesResponse, error) {
	res := service.LanguagesResponse{
		Languages: make([]string, 0, len(languages)),
	}

	for language := range languages {
		res.Languages = append(res.Languages, language)
	}

	return &res, nil
}

var _ service.CodyglotServer = (*Server)(nil)

func initLanguages() error {
	languages = make(map[string]string)

	for _, executor := range config.Executors {
		executorLanguages, err := getExecutorLanguages(executor)
		if err != nil {
			return errors.Wrap(err, "Router: Could not retrieve executor languages")
		}

		for _, language := range executorLanguages {
			languages[language] = executor
		}
	}

	return nil
}

func getExecutorLanguages(executor string) ([]string, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(executor, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Router: could not create executor dial context")
	}
	defer conn.Close()

	client := service.NewCodyglotClient(conn)

	res, err := client.Languages(context.Background(), &service.LanguagesRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Router: an error occured while calling executor")
	}

	return res.Languages, nil
}
