package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/spf13/cobra"
)

// Cmd is the router command group
var Cmd = &cobra.Command{
	Short: "Codyglot router commands",
	Use:   "router",
}

func init() {
	cmd.Cmd.AddCommand(Cmd)
}
