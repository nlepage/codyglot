package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Zenika/codyglot/router/service"

	"google.golang.org/grpc"
)

const (
	defaultPort = 8080
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", defaultPort, "Listening port, default 8080")
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service.RegisterRouterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type server struct{}

func (*server) Execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	return nil, nil
}

var _ service.RouterServer = (*server)(nil)
