package get

import (
	"context"

	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/filestore/client"
	"github.com/nlepage/codyglot/filestore/client/get/config"
	service "github.com/nlepage/codyglot/service/filestore"
)

func Get(id string) error {
	return client.GetClient(func(client service.FileStoreClient) error {
		// FIXME wrap errors

		req, err := client.Get(context.Background(), &service.Id{Id: id})
		if err != nil {
			return err
		}

		if err := filestore.Write(req, config.Config.OutputDir); err != nil {
			return err
		}

		return req.CloseSend()
	}, config.Config.FileStoreClientConfig)
}
