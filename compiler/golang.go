package compiler

import (
	"context"

	"github.com/nlepage/codyglot/exec"
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/ioutil"
	svc "github.com/nlepage/codyglot/service/compiler"
	fssvc "github.com/nlepage/codyglot/service/filestore"
	"github.com/pkg/errors"
)

type Golang Config

func (config Golang) Compiler() *Compiler {
	return &Compiler{config.compile, Config(config)}
}

func (config Golang) compile(ctx context.Context, srcId *fssvc.Id) (*svc.CompileResult, error) {
	tmpDir, err := ioutil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	srcDir := tmpDir.Join("src")

	if err := filestore.Get(srcId.Id, srcDir, config.FilestoreConfig); err != nil {
		return nil, err
	}

	binFile := tmpDir.Join("main")

	buildCmd := exec.Command(ctx, "go", "build", "-o", binFile, srcDir).WithDir(tmpDir.Path())

	if err = buildCmd.Run(); err != nil {
		return nil, errors.Wrap(err, "execute: Build command failed")
	}

	buildRes := buildCmd.CommandResult()

	var binID *fssvc.Id

	// FIXME refactor Put to return id
	if buildRes.Status == 0 {
		err = filestore.Put([]string{binFile}, config.FilestoreConfig)
	}

	return &svc.CompileResult{
		FileStoreId: binID,
		Result:      buildRes,
	}, nil
}
