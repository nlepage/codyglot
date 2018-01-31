package cmd

import (
	"github.com/nlepage/codyglot/executor/cmd"
	"github.com/nlepage/codyglot/executor/golang"
	"github.com/spf13/cobra"
)

var _cmd = &cobra.Command{
	Use:   "golang",
	Short: "Start Codyglot Go(lang) executor",
	RunE: func(_ *cobra.Command, _ []string) error {
		return golang.Executor().Serve()
	},
}

func init() {
	cmd.Cmd.AddCommand(_cmd)
}
