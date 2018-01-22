package golang

import (
	"github.com/Zenika/codyglot/cmd/executor"
	"github.com/spf13/cobra"
)

const (
	defaultPort = 9090
)

var (
	port int
)

var cmd = &cobra.Command{
	Use:   "golang",
	Short: "Start Codyglot golang executor",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	cmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Listening port")
	executor.Cmd.AddCommand(cmd)
}
