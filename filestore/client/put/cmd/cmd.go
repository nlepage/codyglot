package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/filestore/client/cmd"
	"github.com/nlepage/codyglot/filestore/client/put/config"
	"github.com/spf13/cobra"
)

var _cmd = &cobra.Command{
	Short: "Put file(s) to file store server",
	Use:   "put",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		return filestore.Put(args)
	},
}

func init() {
	cmd.AddCommand(_cmd, &config.Config)
}
