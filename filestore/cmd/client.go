package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

func addClientCommand(_cmd *cobra.Command, config *filestore.ClientConfig) {
	_cmd.Flags().StringVarP(&config.Host, "host", "H", filestore.DefaultHost, "File store server host")
	_cmd.Flags().IntVarP(&config.Port, "port", "p", filestore.DefaultPort, "File store server port")
	addCommand(_cmd, &config.Config)
}
