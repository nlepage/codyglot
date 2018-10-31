package cmd

import (
	"github.com/nlepage/codyglot/filestore/client/config"
	"github.com/nlepage/codyglot/filestore/cmd"
	"github.com/spf13/cobra"
)

// AddCommand adds a subcommand to filestore command group, with client config flags
func AddCommand(_cmd *cobra.Command, cfg *config.FileStoreClientConfig) {
	_cmd.Flags().StringVarP(&cfg.Host, "host", "H", config.DefaultHost, "File store server host")
	_cmd.Flags().IntVarP(&cfg.Port, "port", "p", config.DefaultPort, "File store server port")
	cmd.AddCommand(_cmd, &cfg.FileStoreConfig)
}
