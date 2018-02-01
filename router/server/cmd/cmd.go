package cmd

import (
	"github.com/nlepage/codyglot/router/cmd"
	"github.com/nlepage/codyglot/router/server"
	"github.com/nlepage/codyglot/router/server/config"
	"github.com/spf13/cobra"
)

var _cmd = &cobra.Command{
	Use:   "server",
	Short: "Start Codyglot router gRPC server",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		return server.Init()
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		return (&server.Server{}).Serve()
	},
}

func init() {
	_cmd.Flags().IntVarP(&config.Port, "port", "p", config.DefaultPort, "Listening port")
	_cmd.Flags().StringSliceVarP(&config.Languages, "language", "l", nil, "Language")
	cmd.Cmd.AddCommand(_cmd)
}