package gateway

import (
	"github.com/nlepage/codyglot/cmd/router"
	config "github.com/nlepage/codyglot/config/router/gateway"
	"github.com/nlepage/codyglot/router/gateway"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start Codyglot router REST gateway",
	RunE: func(cmd *cobra.Command, args []string) error {
		return (&gateway.Gateway{}).Serve()
	},
}

func init() {
	cmd.Flags().IntVarP(&config.Port, "port", "p", config.DefaultPort, "Listening port")
	cmd.Flags().StringVarP(&config.Endpoint, "endpoint", "e", config.DefaultEndpoint, "gRPC endpoint")
	router.Cmd.AddCommand(cmd)
}
