package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/filestore/cmd"
	"github.com/spf13/cobra"
)

var (
	server = &filestore.Server{}

	_cmd = &cobra.Command{
		Short: "Starts a file store server",
		Use:   "server",
		RunE: func(_ *cobra.Command, _ []string) error {
			return server.Serve()
		},
		PreRunE: func(_ *cobra.Command, _ []string) error {
			return server.Init()
		},
	}
)

func init() {
	_cmd.Flags().IntVarP(&server.Config.Port, "port", "p", filestore.DefaultPort, "Listening port")
	_cmd.Flags().StringVarP(&server.Config.Root, "root", "d", filestore.DefaultRoot, "Root directory")
	cmd.AddCommand(_cmd, &server.Config.Config)
}
