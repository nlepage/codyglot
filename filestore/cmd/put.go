package cmd

import (
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
			return filestore.Put(args, putConfig)
		},
	}
)

func init() {
	addClientCommand(_cmd, &putConfig)
}
