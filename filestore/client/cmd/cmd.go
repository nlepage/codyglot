package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/filestore/cmd"
	"github.com/spf13/cobra"
)

// AddCommand adds a subcommand to filestore command group, with client config flags
func AddCommand(_cmd *cobra.Command, config *filestore.ClientConfig) {
	_cmd.Flags().StringVarP(&config.Host, "host", "H", filestore.DefaultHost, "File store server host")
	_cmd.Flags().IntVarP(&config.Port, "port", "p", filestore.DefaultPort, "File store server port")
	cmd.AddCommand(_cmd, &config.Config)
}
