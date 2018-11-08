package golang

import (
	"context"

	"github.com/nlepage/codyglot/exec"
	"github.com/nlepage/codyglot/executor"
	"github.com/nlepage/codyglot/ioutil"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
)

const (
	language = "golang"
)

// Executor returns Golang Executor
func Executor() *executor.Executor {
	return executor.New(execute, []string{language})
}

// Execute executes Go(lang) code
func execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	if req.Language != language {
		return nil, errors.Errorf("execute: Unsupported language %s", req.Language)
	}

	tmpDir, err := ioutil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	if err = tmpDir.WriteSources(req.Sources); err != nil {
		return nil, err
	}

	binFile := tmpDir.Join("main")

	buildCmd := exec.Command(ctx, "go", "build", "-o", binFile, ".").WithDir(tmpDir.Path())

	if err = buildCmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Build command failed")
	}

	buildRes := buildCmd.CommandResult()

	if buildRes.Status != 0 {
		return &service.ExecuteResponse{
			Compilation: buildRes,
		}, nil
	}

	execRes := make([]*service.CommandResult, 0, len(req.Executions))

	for _, execReq := range req.Executions {
		runCmd := exec.Command(ctx, binFile).WithStdin(execReq.Stdin)

		if err := runCmd.Run(); err != nil {
			return nil, errors.Wrap(err, "execute: Run command failed")
		}

		execRes = append(execRes, runCmd.CommandResult())
	}

	return &service.ExecuteResponse{
		Compilation: buildRes,
		Executions:  execRes,
	}, nil
}
