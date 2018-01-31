package golang

import (
	"context"

	"github.com/nlepage/codyglot/executor"
	"github.com/nlepage/codyglot/executor/executil"
	"github.com/nlepage/codyglot/executor/service"
	"github.com/nlepage/codyglot/executor/tmputil"
	"github.com/pkg/errors"
)

// Executor returns Golang Executor
func Executor() executor.Executor {
	return executor.Executor(execute)
}

// Execute executes Go(lang) code
func execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	tmpDir, err := tmputil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	srcFile, err := tmpDir.WriteFile("main.go", req.Source)
	if err != nil {
		return nil, errors.Wrap(err, "execute: Failed to write source file")
	}

	binFile := tmpDir.Join("main")

	buildCmd := executil.Command(ctx, "go", "build", "-o", binFile, srcFile)

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
		CompilationTime: int64(buildCmd.Duration()),
		RunningTime:     int64(runCmd.Duration()),
	}, nil
}
