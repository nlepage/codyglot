package cmd

import (
	"fmt"

	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

var (
	putConfig filestore.ClientConfig

	_cmd = &cobra.Command{
		Short: "Put file(s) to file store server",
		Use:   "put",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			id, err := filestore.Put(args, putConfig)
			if err != nil {
				return err
			}

			fmt.Println(id.Id)

			return nil
		},
	}
)

func init() {
	addClientCommand(_cmd, &putConfig)
}
