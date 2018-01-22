package gateway

import (
	"context"
	"fmt"
	"net/http"

	router "github.com/Zenika/codyglot/router/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Gateway is Codyglot router gateway
type Gateway struct {
	Port     int
	Endpoint string
}

// Serve starts listening for HTTP requests
func (gw *Gateway) Serve() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := router.RegisterRouterHandlerFromEndpoint(ctx, mux, gw.Endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", gw.Port), mux)
}
