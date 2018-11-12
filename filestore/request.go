package filestore

import (
	"fmt"

	service "github.com/nlepage/codyglot/service/filestore"
	"google.golang.org/grpc"
)

func request(fn func(service.FileStoreClient) error, cfg ClientConfig) error {
	hostport := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(hostport, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	return fn(service.NewFileStoreClient(conn))
}
