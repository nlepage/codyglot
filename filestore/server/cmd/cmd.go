package cmd

import (
	"github.com/nlepage/codyglot/filestore/cmd"
	"github.com/nlepage/codyglot/filestore/server"
	"github.com/nlepage/codyglot/filestore/server/config"
	"github.com/spf13/cobra"
)

var _cmd = &cobra.Command{
	Short: "Starts a file store server",
	Use:   "server",
	RunE: func(_ *cobra.Command, _ []string) error {
		return (&server.Server{}).Serve()
	},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return server.Init()
	},
}

func init() {
	_cmd.Flags().IntVarP(&config.Config.Port, "port", "p", config.DefaultPort, "Listening port")
	_cmd.Flags().StringVarP(&config.Config.Root, "root", "d", config.DefaultRoot, "Root directory")
	cmd.AddCommand(_cmd, &config.Config.FileStoreConfig)
}
