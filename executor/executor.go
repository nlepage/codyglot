package executor

import (
	"context"
	"fmt"
	"net"

	"github.com/nlepage/codyglot/executor/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Executor is a function implementing ExecutorServer
type Executor func(context.Context, *service.ExecuteRequest) (*service.ExecuteResponse, error)

// Execute calls Executor
func (e Executor) Execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	return e(ctx, req)
}

var _ service.ExecutorServer = Executor(nil)

// Serve starts listening for gRPC requests
func (e Executor) Serve(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "Executor: Failed to listen")
	}

	grpcSrv := grpc.NewServer()
	service.RegisterExecutorServer(grpcSrv, e)
	if err := grpcSrv.Serve(lis); err != nil {
		return errors.Wrap(err, "Executor: Failed to serve")
	}

	return nil
}