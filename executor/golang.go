package executor

import (
	"context"

	"github.com/nlepage/codyglot/compiler"
	"github.com/nlepage/codyglot/exec"
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/ioutil"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
)

const golang = "golang"

type Golang Config

func (config Golang) Executor() *Executor {
	return New(Config(config), config.execute, []string{golang})
}

func (config Golang) execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	if req.Language != golang {
		return nil, errors.Errorf("execute: Unsupported language %s", req.Language)
	}

	tmpDir, err := ioutil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	if err = ioutil.WriteSources(tmpDir.Path(), req.Sources); err != nil {
		return nil, err
	}

	srcID, err := filestore.Put(filestore.FsReader([]string{tmpDir.Path()}, config.Filestore.Config, false), config.Filestore)
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

	if err := filestore.Get(buildRes.FileStoreId.Id, filestore.FsWriter(tmpDir.Path()), config.Filestore); err != nil {
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
