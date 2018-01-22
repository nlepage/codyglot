package golang

import (
	"context"
	"fmt"
	"net"

	executor "github.com/Zenika/codyglot/executor/service"
	"google.golang.org/grpc"
)

// Executor is Codyglot Go(lang) executor
type Executor struct {
	Port int
}

// Serve starts listening for gRPC requests
func (e *Executor) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", e.Port))
	if err != nil {
		return fmt.Errorf("Failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	executor.RegisterExecutorServer(grpcSrv, e)
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v", err)
	}

	return nil
}

// Execute executes Go(lang) code
func (e *Executor) Execute(ctx context.Context, req *executor.ExecuteRequest) (*executor.ExecuteResponse, error) {
	return &executor.ExecuteResponse{
		ExitCode: 1,
		StdOut:   "golang",
		StdErr:   "",
	}, nil
}

var _ executor.ExecutorServer = (*Executor)(nil)
