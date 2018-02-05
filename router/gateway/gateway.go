package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nlepage/codyglot/router/gateway/config"
	"github.com/nlepage/codyglot/service"
	"google.golang.org/grpc"
)

// Gateway is Codyglot router gateway
type Gateway struct{}

// Serve starts listening for HTTP requests
func (gw *Gateway) Serve() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := service.RegisterCodyglotHandlerFromEndpoint(ctx, mux, config.Endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux)
}
