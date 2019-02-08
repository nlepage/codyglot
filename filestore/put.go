package filestore

import (
	"context"

	service "github.com/nlepage/codyglot/service/filestore"
	log "github.com/sirupsen/logrus"
)

// FIXME log should log on stderr

// Put puts files to a store server
func Put(paths []string, config ClientConfig) (*service.Id, error) {
	var id *service.Id

	err := request(func(client service.FileStoreClient) error {
		// FIXME wrap errors

		req, err := client.Put(context.Background())
		if err != nil {
			return err
		}

		for _, path := range paths {
			if err := send(req, FsReader(path, config.Config, true)); err != nil {
				log.WithError(err).Errorf("Error while walking putting path %s", path)
			}
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
