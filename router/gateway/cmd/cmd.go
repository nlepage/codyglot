package cmd

import (
	"github.com/nlepage/codyglot/router/cmd"
	"github.com/nlepage/codyglot/router/gateway"
	"github.com/nlepage/codyglot/router/gateway/config"
	"github.com/spf13/cobra"
)

var _cmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start Codyglot router REST gateway",
	RunE: func(_ *cobra.Command, _ []string) error {
		return (&gateway.Gateway{}).Serve()
	},
}

func init() {
	_cmd.Flags().IntVarP(&config.Port, "port", "p", config.DefaultPort, "Listening port")
	_cmd.Flags().StringVarP(&config.Endpoint, "endpoint", "e", config.DefaultEndpoint, "gRPC endpoint")
	cmd.Cmd.AddCommand(_cmd)
}
