package put

import (
	"context"
	"fmt"
	"os"

	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/filestore/client"
	"github.com/nlepage/codyglot/filestore/client/put/config"
	service "github.com/nlepage/codyglot/service/filestore"
	log "github.com/sirupsen/logrus"
)

// FIXME log should log on stderr

// Put puts files to a store server
func Put(files []string) error {
	return client.GetClient(func(client service.FileStoreClient) error {
		// FIXME wrap errors

		req, err := client.Put(context.Background())
		if err != nil {
			return err
		}

		for _, file := range files {
			info, err := os.Stat(file)
			if err != nil {
				log.WithError(err).Errorf("Could not determine file type of %s", file)
			}

			if info.IsDir() {
				if err := filestore.SendDir(req, file, config.Config.FileStoreConfig, true); err != nil {
					log.WithError(err).Errorf("Error while walking dir %s", file)
				}
			} else {
				if err := filestore.SendFile(req, file, info.Name(), info, config.Config.FileStoreConfig); err != nil {
					log.WithError(err).Errorf("Could not send file %s", file)
				}
			}
		}

		res, err := req.CloseAndRecv()
		if err != nil {
			return err
		}

		fmt.Println(res.Id)

		return nil
	}, config.Config)
}
