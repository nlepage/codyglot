package filestore

import (
	"context"

	service "github.com/nlepage/codyglot/service/filestore"
)

// FIXME log should log on stderr

// Put puts files to a store server
func Put(r Reader, config ClientConfig) (*service.Id, error) {
	var id *service.Id

	err := request(func(client service.FileStoreClient) error {
		// FIXME wrap errors

		req, err := client.Put(context.Background())
		if err != nil {
			return err
		}

		if err := send(req, r); err != nil {
			return err
		}

		res, err := req.CloseAndRecv()
		if err != nil {
			return err
		}

		id = res

		return nil
	}, config)

	return id, err
}
