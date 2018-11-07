package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/filestore/client/cmd"
	"github.com/nlepage/codyglot/filestore/client/get/config"
	"github.com/spf13/cobra"
)

var _cmd = &cobra.Command{
	Short: "get file(s) from file store server",
	Use:   "get",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		return filestore.Get(args[0])
	},
}

func init() {
	_cmd.Flags().StringVarP(&config.Config.OutputDir, "output-dir", "o", config.DefaultOutputDir, "Output directory")
	cmd.AddCommand(_cmd, &config.Config.FileStoreClientConfig)
}
