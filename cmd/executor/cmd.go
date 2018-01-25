package executor

import (
	"github.com/nlepage/codyglot/cmd/codyglot"
	"github.com/spf13/cobra"
)

const (
	defaultPort = 9090
)

var (
	// Port is the executor listening port
	Port int
)

// Cmd is the executor command group
var Cmd = &cobra.Command{
	Short: "Codyglot executor commands",
	Use:   "executor",
}

func init() {
	Cmd.PersistentFlags().IntVarP(&Port, "port", "p", defaultPort, "Listening port")
	codyglot.Cmd.AddCommand(Cmd)
}
