package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/compiler/config"
	"github.com/spf13/cobra"
)

// Cmd is the compiler command group
var Cmd = &cobra.Command{
	Short: "Codyglot compiler commands",
	Use:   "compiler",
}

func init() {
	Cmd.PersistentFlags().IntVarP(&config.Config.Port, "port", "p", config.DefaultPort, "Listening port")
	cmd.Cmd.AddCommand(Cmd)
}
