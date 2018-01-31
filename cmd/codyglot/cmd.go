package codyglot

import (
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
}
