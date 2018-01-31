package golang

import (
	"github.com/nlepage/codyglot/cmd/executor"
	"github.com/nlepage/codyglot/executor/golang"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "golang",
	Short: "Start Codyglot Go(lang) executor",
	RunE: func(cmd *cobra.Command, args []string) error {
		return golang.Executor().Serve()
	},
}

func init() {
	executor.Cmd.AddCommand(cmd)
}
