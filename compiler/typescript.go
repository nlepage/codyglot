package compiler

import (
	"context"

	"github.com/nlepage/codyglot/exec"
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/ioutil"
	svc "github.com/nlepage/codyglot/service/compiler"
	fssvc "github.com/nlepage/codyglot/service/filestore"
)

type TypeScript ServerConfig

func (config TypeScript) Server() *Server {
	return &Server{config.compile, ServerConfig(config)}
}

func (config TypeScript) compile(ctx context.Context, srcId *fssvc.Id) (*svc.CompileResult, error) {
	// FIXME errors

	tmpDir, err := ioutil.NewTmpDir()
	if err != nil {
		return nil, err
	}
	defer tmpDir.Close()

	srcDir := tmpDir.Join("src")

	if err := filestore.Get(srcId.Id, filestore.FsWriter(srcDir), config.Filestore); err != nil {
		return nil, err
	}

	initCmd := exec.Command(ctx, "tsc", "--init").WithDir(srcDir)

	if err = initCmd.Run(); err != nil {
		return nil, err
	}

	if initRes := initCmd.CommandResult(); initRes.Status != 0 {
		return &svc.CompileResult{
			Result: initRes,
		}, nil
	}

	outDir := tmpDir.Join("out")

	compileCmd := exec.Command(ctx, "tsc", "--outDir", outDir).WithDir(srcDir)

	if err = compileCmd.Run(); err != nil {
		return nil, err
	}

	compileRes := compileCmd.CommandResult()

	if compileRes.Status != 0 {
		return &svc.CompileResult{
			Result: compileRes,
		}, nil
	}

	outID, err := filestore.Put(filestore.FsReader([]string{outDir}, config.Filestore.Config, false), config.Filestore)
	if err != nil {
		return nil, err
	}

	return &svc.CompileResult{
		FileStoreId: outID,
		Result:      compileRes,
	}, nil
}
