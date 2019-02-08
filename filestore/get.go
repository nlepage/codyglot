package filestore

import (
	"context"

	service "github.com/nlepage/codyglot/service/filestore"
)

func Get(id string, w Writer, config ClientConfig) error {
	return request(func(client service.FileStoreClient) error {
		// FIXME wrap errors

		req, err := client.Get(context.Background(), &service.Id{Id: id})
		if err != nil {
			return err
		}

		if err := recv(req, w); err != nil {
			return err
		}

		return req.CloseSend()
	}, config)
}
