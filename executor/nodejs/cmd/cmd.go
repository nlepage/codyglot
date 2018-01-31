package cmd

import (
	"github.com/nlepage/codyglot/executor/cmd"
	"github.com/nlepage/codyglot/executor/nodejs"
	"github.com/spf13/cobra"
)

var _cmd = &cobra.Command{
	Use:   "nodejs",
	Short: "Start Codyglot NodeJS executor",
	RunE: func(_ *cobra.Command, _ []string) error {
		return nodejs.Executor().Serve()
	},
}

func init() {
	cmd.Cmd.AddCommand(_cmd)
}
