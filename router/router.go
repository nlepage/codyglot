package main

import (
	"context"
	"fmt"
	"log"
	"net"

	router "github.com/Zenika/codyglot/router/service"
	"github.com/spf13/cobra"

	"google.golang.org/grpc"
)

const (
	defaultPort = 8080
)

var (
	port int
)

func init() {
	cmd.PersistentFlags().IntVarP(&port, "port", "p", defaultPort, "Listening port, default 8080")
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var cmd = &cobra.Command{
	Use: "router",
	RunE: func(cmd *cobra.Command, args []string) error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return fmt.Errorf("Failed to listen: %v", err)
		}

		s := grpc.NewServer()
		router.RegisterRouterServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			return fmt.Errorf("Failed to serve: %v", err)
		}
		return nil
	},
}

type server struct{}

func (*server) Execute(ctx context.Context, req *router.ExecuteRequest) (*router.ExecuteResponse, error) {
	return nil, nil
}

var _ router.RouterServer = (*server)(nil)
