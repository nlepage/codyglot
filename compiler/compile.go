package compiler

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	svc "github.com/nlepage/codyglot/service/compiler"
	fssvc "github.com/nlepage/codyglot/service/filestore"
)

func Compile(ctx context.Context, srcId *fssvc.Id, config ClientConfig) (*svc.CompileResult, error) {
	// FIXME errors

	hostport := fmt.Sprintf("%s:%d", config.Host, config.Port)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(hostport, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := svc.NewCompilerClient(conn)

	return client.Compile(ctx, srcId)
}
