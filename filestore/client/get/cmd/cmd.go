package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/filestore/client/cmd"
	"github.com/spf13/cobra"
)

var (
	config filestore.GetConfig

	_cmd = &cobra.Command{
		Short: "get file(s) from file store server",
		Use:   "get",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return filestore.Get(args[0], config)
		},
	}
)

func init() {
	_cmd.Flags().StringVarP(&config.OutputDir, "output-dir", "o", filestore.DefaultOutputDir, "Output directory")
	cmd.AddCommand(_cmd, &config.ClientConfig)
}
