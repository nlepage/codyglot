package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

var (
	server = &filestore.Server{}

	serverCmd = &cobra.Command{
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
	serverCmd.Flags().IntVarP(&server.Config.Port, "port", "p", filestore.DefaultPort, "Listening port")
	serverCmd.Flags().StringVarP(&server.Config.Root, "root", "d", filestore.DefaultRoot, "Root directory")
	addCommand(serverCmd, &server.Config.Config)
}
