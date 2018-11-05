package golang

import (
	"context"

	"github.com/nlepage/codyglot/compiler"
	svc "github.com/nlepage/codyglot/service/compiler"
	fssvc "github.com/nlepage/codyglot/service/filestore"
)

var Compiler = compiler.New(compile)

func compile(context.Context, *fssvc.Id) (*svc.CompileResult, error) {
	return &svc.CompileResult{}, nil
}
