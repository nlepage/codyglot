package executor

import (
	"github.com/nlepage/codyglot/cmd/codyglot"
	config "github.com/nlepage/codyglot/config/executor"
	"github.com/nlepage/codyglot/executor"
	"github.com/spf13/cobra"
)

// Cmd is the executor command group
var Cmd = &cobra.Command{
	Short: "Codyglot executor commands",
	Use:   "executor",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		executor.Init()
	},
}

func init() {
	Cmd.PersistentFlags().IntVarP(&config.Port, "port", "p", config.DefaultPort, "Listening port")
	Cmd.PersistentFlags().IntVar(&config.CleanupBuffer, "cleanup-buffer", config.DefaultCleanupBuffer, "Size of the cleanup buffer")
	Cmd.PersistentFlags().IntVar(&config.CleanupRoutines, "cleanup-routines", config.DefaultCleanupRoutines, "Number of cleanup routines")
	codyglot.Cmd.AddCommand(Cmd)
}
