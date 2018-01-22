package server

import (
	"context"
	"fmt"
	"net"

	"github.com/Zenika/codyglot/router/service"
	"google.golang.org/grpc"
)

type Server struct {
	LanguagesMap map[string]string
	Port         int
}

func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return fmt.Errorf("Failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	service.RegisterRouterServer(grpcSrv, s)
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v", err)
	}

	return nil
}

func (s *Server) Execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	return nil, nil
}

func (s *Server) Languages(ctx context.Context, req *service.LanguagesRequest) (*service.LanguagesResponse, error) {
	res := service.LanguagesResponse{
		Languages: make([]string, 0, len(s.LanguagesMap)),
	}

	for language := range s.LanguagesMap {
		res.Languages = append(res.Languages, language)
	}

	return &res, nil
}

var _ service.RouterServer = (*Server)(nil)
