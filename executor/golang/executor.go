package golang

import (
	"context"
	"io/ioutil"
	"os"
	"path"

	"github.com/nlepage/codyglot/executor"
	"github.com/nlepage/codyglot/executor/internal/executil"
	"github.com/nlepage/codyglot/executor/service"
	"github.com/pkg/errors"
)

// Executor returns Golang Executor
func Executor() executor.Executor {
	return executor.Executor(execute)
}

// Execute executes Go(lang) code
func execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	p, err := writeSourceFile(req.Source)
	if err != nil {
		return nil, err
	}

	cmd := executil.Command(ctx, "go", "run", p).WithStdin(req.Stdin)

	if err = cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "GolangExecutor: Failed to run command")
	}

	return &service.ExecuteResponse{
		ExitStatus: cmd.ExitStatus(),
		Stdout:     cmd.Stdout(),
		Stderr:     cmd.Stderr(),
	}, nil
}

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
