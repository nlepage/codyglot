package compiler

import (
	"context"
	"fmt"
	"net"

	svc "github.com/nlepage/codyglot/service/compiler"
	fssvc "github.com/nlepage/codyglot/service/filestore"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Compiler is a struct implementing CompilerServer as a specific compiler
type Compiler struct {
	fn     func(context.Context, *fssvc.Id) (*svc.CompileResult, error)
	config Config
}

// Compile calls Compiler.fn
func (c *Compiler) Compile(ctx context.Context, fsID *fssvc.Id) (*svc.CompileResult, error) {
	return c.fn(ctx, fsID)
}

var _ svc.CompilerServer = (*Compiler)(nil)

// Serve starts listening for gRPC requests
func (c *Compiler) Serve() error {
	log.Infoln("Starting compiler on port", c.config.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.config.Port))
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
