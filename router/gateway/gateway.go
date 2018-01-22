package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Zenika/codyglot/router/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Gateway struct {
	Port     int
	Endpoint string
}

func (gw *Gateway) Serve() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := service.RegisterRouterHandlerFromEndpoint(ctx, mux, gw.Endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", gw.Port), mux)
}
