package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

var (
	getConfig filestore.GetConfig

	getCmd = &cobra.Command{
		Short: "get file(s) from file store server",
		Use:   "get",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return filestore.Get(args[0], getConfig)
		},
	}
)

func init() {
	getCmd.Flags().StringVarP(&getConfig.OutputDir, "output-dir", "o", filestore.DefaultOutputDir, "Output directory")
	addClientCommand(getCmd, &getConfig.ClientConfig)
}
