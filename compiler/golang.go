package compiler

import (
	"context"

	"github.com/nlepage/codyglot/exec"
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/ioutil"
	svc "github.com/nlepage/codyglot/service/compiler"
	fssvc "github.com/nlepage/codyglot/service/filestore"
)

type Golang ServerConfig

func (config Golang) Server() *Server {
	return &Server{config.compile, ServerConfig(config)}
}

func (config Golang) compile(ctx context.Context, srcId *fssvc.Id) (*svc.CompileResult, error) {
	// FIXME errors

	tmpDir, err := ioutil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	srcDir := tmpDir.Join("src")

	if err := filestore.Get(srcId.Id, srcDir, config.Filestore); err != nil {
		return nil, err
	}

	binFile := tmpDir.Join("main")

	buildCmd := exec.Command(ctx, "go", "build", "-o", binFile, srcDir).WithDir(tmpDir.Path())

	if err = buildCmd.Run(); err != nil {
		return nil, err
	}

	buildRes := buildCmd.CommandResult()

	if buildRes.Status != 0 {
		return &svc.CompileResult{
			Result: buildRes,
		}, nil
	}

	binID, err := filestore.Put([]string{binFile}, config.Filestore)
	if err != nil {
		return nil, err
	}

	return &svc.CompileResult{
		FileStoreId: binID,
		Result:      buildRes,
	}, nil
}
