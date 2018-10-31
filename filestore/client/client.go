package client

import (
	"fmt"

	"github.com/nlepage/codyglot/filestore/client/config"
	"github.com/nlepage/codyglot/service/filestore"
	"google.golang.org/grpc"
)

func GetClient(fn func(filestore.FileStoreClient) error, cfg config.FileStoreClientConfig) error {
	hostport := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(hostport, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	return fn(filestore.NewFileStoreClient(conn))
}
