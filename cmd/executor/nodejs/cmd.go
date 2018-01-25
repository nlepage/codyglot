package golang

import (
	"github.com/nlepage/codyglot/cmd/executor"
	"github.com/nlepage/codyglot/executor/nodejs"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "nodejs",
	Short: "Start Codyglot NodeJS executor",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nodejs.Executor().Serve(executor.Port)
	},
}

func init() {
	executor.Cmd.AddCommand(cmd)
}
