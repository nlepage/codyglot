package ping

import (
	"context"

	"github.com/nlepage/codyglot/service"
)

// PingServer implements Ping of Codyglot service
type Server struct{}

// Ping answers pong
func (ps *Server) Ping(ctx context.Context, req *service.Ping) (*service.Pong, error) {
	return &service.Pong{}, nil
}
