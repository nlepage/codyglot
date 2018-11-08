package golang

import (
	"context"
	"github.com/nlepage/codyglot/compiler"
	"github.com/nlepage/codyglot/executor/config"
	"os"

	"github.com/nlepage/codyglot/filestore"

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

	files, err := listFiles(tmpDir)
	if err != nil {
		return nil, err
	}

	srcID, err := filestore.Put(files, config.Filestore)
	if err != nil {
		return nil, err
	}

	buildRes, err := compiler.Compile(ctx, srcID, config.Compiler)
	if err != nil {
		return nil, err
	}

	if buildRes.Result.Status != 0 {
		return &service.ExecuteResponse{
			Compilation: buildRes.Result,
		}, nil
	}

	if err := filestore.Get(buildRes.FileStoreId.Id, tmpDir.Path(), config.Filestore); err != nil {
		return nil, err
	}

	execRes := make([]*service.CommandResult, 0, len(req.Executions))

	for _, execReq := range req.Executions {
		runCmd := exec.Command(ctx, tmpDir.Join("main")).WithStdin(execReq.Stdin)

		if err := runCmd.Run(); err != nil {
			return nil, errors.Wrap(err, "execute: Run command failed")
		}

		execRes = append(execRes, runCmd.CommandResult())
	}

	return &service.ExecuteResponse{
		Compilation: buildRes.Result,
		Executions:  execRes,
	}, nil
}

func listFiles(tmpDir *ioutil.TmpDir) ([]string, error) {
	f, err := os.Open(tmpDir.Path())
	if err != nil {
		return nil, err
	}

	files, err := f.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	for i, f := range files {
		files[i] = tmpDir.Join(f)
	}

	return  files, nil
}
