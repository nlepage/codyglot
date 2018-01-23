package gateway

import (
	"github.com/nlepage/codyglot/cmd/router"
	"github.com/nlepage/codyglot/router/gateway"
	"github.com/spf13/cobra"
)

const (
	defaultPort     = 8080
	defaultEndpoint = "localhost:9090"
)

var (
	port     int
	endpoint string
)

var cmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start Codyglot router REST gateway",
	RunE: func(cmd *cobra.Command, args []string) error {
		gw := &gateway.Gateway{
			Port:     port,
			Endpoint: endpoint,
		}

		return gw.Serve()
	},
}

func init() {
	cmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Listening port")
	cmd.Flags().StringVarP(&endpoint, "endpoint", "e", defaultEndpoint, "gRPC endpoint")
	router.Cmd.AddCommand(cmd)
}
