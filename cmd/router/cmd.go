package router

import (
	"github.com/Zenika/codyglot/cmd/codyglot"
	"github.com/spf13/cobra"
)

// Cmd is the router command group
var Cmd = &cobra.Command{
	Short: "Codyglot router commands",
	Use:   "router",
}

func init() {
	codyglot.Cmd.AddCommand(Cmd)
}
