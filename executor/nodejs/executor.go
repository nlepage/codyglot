package nodejs

import (
	"context"

	"github.com/nlepage/codyglot/executor"
	"github.com/nlepage/codyglot/executor/internal/executil"
	"github.com/nlepage/codyglot/executor/internal/srcutil"
	"github.com/nlepage/codyglot/executor/service"
	"github.com/pkg/errors"
)

// Executor returns NodeJS Executor
func Executor() executor.Executor {
	return executor.Executor(execute)
}

// Execute executes NodeJS code
func execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	p, err := srcutil.WriteSourceFile("index.js", req.Source)
	if err != nil {
		return nil, err
	}

	cmd := executil.Command(ctx, "node", p).WithStdin(req.Stdin)

	if err = cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Failed to run command")
	}

	return &service.ExecuteResponse{
		ExitStatus: cmd.ExitStatus(),
		Stdout:     cmd.Stdout(),
		Stderr:     cmd.Stderr(),
	}, nil
}