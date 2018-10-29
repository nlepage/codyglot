package golang

import (
	"context"

	"github.com/nlepage/codyglot/executor"
	"github.com/nlepage/codyglot/executor/executil"
	"github.com/nlepage/codyglot/executor/tmputil"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
)

const (
	golang = "golang"
)

// Executor returns Golang Executor
func Executor() *executor.Executor {
	return executor.New(execute, []string{golang})
}

// Execute executes Go(lang) code
func execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	if req.Language != golang {
		return nil, errors.Errorf("execute: Unsupported language %s", req.Language)
	}

	tmpDir, err := tmputil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	if err = tmpDir.WriteSources(req.Sources); err != nil {
		return nil, err
	}

	binFile := tmpDir.Join("main")

	buildCmd := executil.Command(ctx, "go", "build", "-o", binFile, ".").WithDir(tmpDir.Path())

	if err = buildCmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Build command failed")
	}

	if buildCmd.ExitStatus() != 0 {
		return &service.ExecuteResponse{
			ExitStatus: buildCmd.ExitStatus(),
			Stdout:     buildCmd.Stdout(),
			Stderr:     buildCmd.Stderr(),
		}, nil
	}

	runCmd := executil.Command(ctx, binFile).WithStdin(req.Stdin)

	if err := runCmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Run command failed")
	}

	return &service.ExecuteResponse{
		ExitStatus:      runCmd.ExitStatus(),
		Stdout:          runCmd.Stdout(),
		Stderr:          runCmd.Stderr(),
		CompilationTime: buildCmd.Duration(),
		RunningTime:     runCmd.Duration(),
	}, nil
}
