package compiler

import (
	"context"
	"fmt"
	"net"

	"github.com/nlepage/codyglot/compiler/config"
	svc "github.com/nlepage/codyglot/service/compiler"
	fssvc "github.com/nlepage/codyglot/service/filestore"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Compiler is a struct implementing CompilerServer as a specific compiler
type Compiler struct {
	// TODO remove struct ?
	fn func(context.Context, *fssvc.Id) (*svc.CompileResult, error)
}

// New creates a new Compiler
func New(fn func(context.Context, *fssvc.Id) (*svc.CompileResult, error)) *Compiler {
	return &Compiler{fn}
}

// Compile calls Compiler.fn
func (c *Compiler) Compile(ctx context.Context, fsID *fssvc.Id) (*svc.CompileResult, error) {
	return c.fn(ctx, fsID)
}

var _ svc.CompilerServer = (*Compiler)(nil)

// Serve starts listening for gRPC requests
func (c *Compiler) Serve() error {
	log.Infoln("Starting compiler on port", config.Config.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Config.Port))
	if err != nil {
		return errors.Wrap(err, "Compiler: Failed to listen")
	}

	grpcSrv := grpc.NewServer()
	svc.RegisterCompilerServer(grpcSrv, c)
	if err := grpcSrv.Serve(lis); err != nil {
		return errors.Wrap(err, "Compiler: Failed to serve")
	}

	return nil
}
