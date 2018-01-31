package server

import (
	"github.com/nlepage/codyglot/cmd/router"
	"github.com/nlepage/codyglot/config/router/server"
	"github.com/nlepage/codyglot/router/server"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "server",
	Short: "Start Codyglot router gRPC server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return server.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return (&server.Server{}).Serve()
	},
}

func init() {
	cmd.Flags().IntVarP(&config.Port, "port", "p", config.DefaultPort, "Listening port")
	cmd.Flags().StringSliceVarP(&config.Languages, "language", "l", nil, "Language")
	router.Cmd.AddCommand(cmd)
}
