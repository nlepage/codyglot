package filestore

import (
	"context"

	service "github.com/nlepage/codyglot/service/filestore"
)

func Get(id string, config GetConfig) error {
	return request(func(client service.FileStoreClient) error {
		// FIXME wrap errors

		req, err := client.Get(context.Background(), &service.Id{Id: id})
		if err != nil {
			return err
		}

		if err := recv(req, config.OutputDir); err != nil {
			return err
		}

		return req.CloseSend()
	}, config.ClientConfig)
}
