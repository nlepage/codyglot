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

func executeJavascript(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	tmpDir, err := tmputil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	srcFile, err := tmpDir.WriteFile("index.js", req.Source)
	if err != nil {
		return nil, err
	}

	cmd := executil.Command(ctx, "node", srcFile).WithStdin(req.Stdin)

	if err = cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Failed to run command")
	}

	return &service.ExecuteResponse{
		ExitStatus:  cmd.ExitStatus(),
		Stdout:      cmd.Stdout(),
		Stderr:      cmd.Stderr(),
		RunningTime: cmd.Duration(),
	}, nil
}

func executeTypescript(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	tmpDir, err := tmputil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	srcFile, err := tmpDir.WriteFile("index.ts", req.Source)
	if err != nil {
		return nil, err
	}

	compileCmd := executil.Command(ctx, "tsc", srcFile)

	if err = compileCmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Compile command failed")
	}

	if compileCmd.ExitStatus() != 0 {
		return &service.ExecuteResponse{
			ExitStatus: compileCmd.ExitStatus(),
			Stdout:     compileCmd.Stdout(),
			Stderr:     compileCmd.Stderr(),
		}, nil
	}

	runCmd := executil.Command(ctx, "node", tmpDir.Join("index.js")).WithStdin(req.Stdin)

	if err = runCmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Run command failed")
	}

	return &service.ExecuteResponse{
		ExitStatus:      runCmd.ExitStatus(),
		Stdout:          runCmd.Stdout(),
		Stderr:          runCmd.Stderr(),
		CompilationTime: compileCmd.Duration(),
		RunningTime:     runCmd.Duration(),
	}, nil
}
