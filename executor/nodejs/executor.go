package nodejs

import (
	"context"

	"github.com/nlepage/codyglot/executor"
	"github.com/nlepage/codyglot/executor/executil"
	"github.com/nlepage/codyglot/executor/tmputil"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
)

const (
	javascript = "javascript"
	typescript = "typescript"
)

// Executor returns NodeJS Executor
func Executor() *executor.Executor {
	return executor.New(execute, []string{javascript, typescript})
}

func execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	switch req.Language {
	case javascript:
		return executeJavascript(ctx, req)
	case typescript:
		return executeTypescript(ctx, req)
	default:
		return nil, errors.Errorf("execute: Unsupported language %s", req.Language)
	}
}

// FIXME wrap errors

func executeJavascript(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	tmpDir, err := tmputil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	if err = tmpDir.WriteSources(req.Sources); err != nil {
		return nil, err
	}

	execResults := make([]*service.CommandResult, 0, len(req.Executions))

	for _, execReq := range req.Executions {
		cmd := executil.Command(ctx, "node", tmpDir.Path()).WithStdin(execReq.Stdin)

		if err = cmd.Run(); err != nil {
			return nil, errors.Wrap(err, "execute: Failed to run command")
		}

		execResults = append(execResults, cmd.CommandResult())
	}

	return &service.ExecuteResponse{
		Executions: execResults,
	}, nil
}

func executeTypescript(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	tmpDir, err := tmputil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	if err = tmpDir.WriteSources(req.Sources); err != nil {
		return nil, err
	}

	initCmd := executil.Command(ctx, "tsc", "--init").WithDir(tmpDir.Path())

	if err = initCmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Init command failed")
	}

	if initRes := initCmd.CommandResult(); initRes.Status != 0 {
		return &service.ExecuteResponse{
			Compilation: initRes,
		}, nil
	}

	compileCmd := executil.Command(ctx, "tsc").WithDir(tmpDir.Path())

	if err = compileCmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Compile command failed")
	}

	compileRes := compileCmd.CommandResult()

	if compileRes.Status != 0 {
		return &service.ExecuteResponse{
			Compilation: compileRes,
		}, nil
	}

	execRes := make([]*service.CommandResult, 0, len(req.Executions))

	for _, execReq := range req.Executions {
		runCmd := executil.Command(ctx, "node", tmpDir.Path()).WithStdin(execReq.Stdin)

		if err = runCmd.Run(); err != nil {
			return nil, errors.Wrap(err, "execute: Run command failed")
		}

		execRes = append(execRes, runCmd.CommandResult())
	}

	return &service.ExecuteResponse{
		Compilation: compileRes,
		Executions:  execRes,
	}, nil
}
