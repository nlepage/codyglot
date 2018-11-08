package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/executor/config"
	"github.com/spf13/cobra"
)

// Cmd is the executor command group
var Cmd = &cobra.Command{
	Short: "Codyglot executor commands",
	Use:   "executor",
}

func init() {
	Cmd.PersistentFlags().IntVarP(&config.Port, "port", "p", config.DefaultPort, "Listening port")
	cmd.Cmd.AddCommand(Cmd)
}
