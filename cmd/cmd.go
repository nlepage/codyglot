package cmd

import (
	"github.com/nlepage/codyglot/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logLevel = logLevelValue{log.GetLevel()}
)

// Cmd is the root command
var Cmd = &cobra.Command{
	Short: "Codygloy is a polyglot code execution tool",
	Use:   "codyglot",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel.Level))
	},
}

func init() {
	Cmd.PersistentFlags().Var(&logLevel, "log-level", "Log level (panic, fatal, error, warn, info, debug)")
	Cmd.PersistentFlags().IntVar(&config.CleanupBuffer, "tmp-cleanup-buffer", config.DefaultCleanupBuffer, "Size of the cleanup buffer for temporary files")
	Cmd.PersistentFlags().IntVar(&config.CleanupRoutines, "tmp-cleanup-routines", config.DefaultCleanupRoutines, "Number of cleanup routines for temporary files")
}
