package executor

import (
	"context"

	"github.com/nlepage/codyglot/compiler"
	"github.com/nlepage/codyglot/exec"
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/ioutil"
	svc "github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
)

const (
	javascript = "javascript"
	typescript = "typescript"
)

type NodeJS Config

// Executor returns NodeJS Executor
func (config NodeJS) Executor() *Executor {
	return New(Config(config), config.execute, []string{javascript, typescript})
}

func (config NodeJS) execute(ctx context.Context, req *svc.ExecuteRequest) (*svc.ExecuteResponse, error) {
	switch req.Language {
	case javascript:
		return config.executeJavascript(ctx, req)
	case typescript:
		return config.executeTypescript(ctx, req)
	default:
		return nil, errors.Errorf("execute: Unsupported language %s", req.Language)
	}
}

// FIXME wrap errors

func (config NodeJS) executeJavascript(ctx context.Context, req *svc.ExecuteRequest) (*svc.ExecuteResponse, error) {
	tmpDir, err := ioutil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	if err = ioutil.WriteSources(tmpDir.Path(), req.Sources); err != nil {
		return nil, err
	}

	execResults := make([]*svc.CommandResult, 0, len(req.Executions))

	for _, execReq := range req.Executions {
		cmd := exec.Command(ctx, "node", tmpDir.Path()).WithStdin(execReq.Stdin)

		if err = cmd.Run(); err != nil {
			return nil, errors.Wrap(err, "execute: Failed to run command")
		}

		execResults = append(execResults, cmd.CommandResult())
	}

	return &svc.ExecuteResponse{
		Executions: execResults,
	}, nil
}

func (config NodeJS) executeTypescript(ctx context.Context, req *svc.ExecuteRequest) (*svc.ExecuteResponse, error) {
	tmpDir, err := ioutil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	srcDir := tmpDir.Join("src")

	if err = ioutil.WriteSources(srcDir, req.Sources); err != nil {
		return nil, err
	}

	srcFiles, err := ioutil.ListFiles(srcDir)
	if err != nil {
		return nil, err
	}

	srcID, err := filestore.Put(srcFiles, config.Filestore)
	if err != nil {
		return nil, err
	}

	compileRes, err := compiler.Compile(ctx, srcID, config.Compiler)
	if err != nil {
		return nil, err
	}

	if compileRes.Result.Status != 0 {
		return &svc.ExecuteResponse{
			Compilation: compileRes.Result,
		}, nil
	}

	outDir := tmpDir.Join("out")

	if err := filestore.Get(compileRes.FileStoreId.Id, outDir, config.Filestore); err != nil {
		return nil, err
	}

	execRes := make([]*svc.CommandResult, 0, len(req.Executions))

	for _, execReq := range req.Executions {
		runCmd := exec.Command(ctx, "node", outDir).WithStdin(execReq.Stdin)

		if err = runCmd.Run(); err != nil {
			return nil, errors.Wrap(err, "execute: Run command failed")
		}

		execRes = append(execRes, runCmd.CommandResult())
	}

	return &svc.ExecuteResponse{
		Compilation: compileRes.Result,
		Executions:  execRes,
	}, nil
}
