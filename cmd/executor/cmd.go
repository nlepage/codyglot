package executor

import (
	"github.com/nlepage/codyglot/cmd/codyglot"
	"github.com/spf13/cobra"
)

// Cmd is the executor command group
var Cmd = &cobra.Command{
	Short: "Codyglot executor commands",
	Use:   "executor",
}

func init() {
	codyglot.Cmd.AddCommand(Cmd)
}
