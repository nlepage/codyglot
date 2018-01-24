package golang

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"

	"github.com/nlepage/codyglot/executor/executil"
	executor "github.com/nlepage/codyglot/executor/service"
	"github.com/pkg/errors"
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
		return errors.Wrap(err, "Executor: Failed to listen")
	}

	grpcSrv := grpc.NewServer()
	executor.RegisterExecutorServer(grpcSrv, e)
	if err := grpcSrv.Serve(lis); err != nil {
		return errors.Wrap(err, "Executor: Failed to serve")
	}

	return nil
}

// Execute executes Go(lang) code
func (e *Executor) Execute(ctx context.Context, req *executor.ExecuteRequest) (*executor.ExecuteResponse, error) {
	p, err := writeSourceFile(req.Source)
	if err != nil {
		return nil, err
	}

	cmd := executil.Command(ctx, "go", "run", p).WithStdin(req.Stdin)

	if err = cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "Executor: Failed to run command")
	}

	return &executor.ExecuteResponse{
		ExitStatus: cmd.ExitStatus(),
		Stdout:     cmd.Stdout(),
		Stderr:     cmd.Stderr(),
	}, nil
}

var _ executor.ExecutorServer = (*Executor)(nil)

func writeSourceFile(code string) (string, error) {
	dir, err := ioutil.TempDir("", "codyglot")
	if err != nil {
		return "", errors.Wrap(err, "Executor: Failed to create temp dir")
	}

	p := path.Join(dir, "main.go")

	f, err := os.Create(p)
	if err != nil {
		return "", errors.Wrap(err, "Executor: Failed to create source file")
	}
	defer f.Close()

	_, err = f.WriteString(code)
	if err != nil {
		return "", errors.Wrap(err, "Executor: Failed to write source file")
	}

	return p, nil
}
