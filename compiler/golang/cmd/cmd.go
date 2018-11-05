package cmd

import (
	"github.com/nlepage/codyglot/compiler/cmd"
	"github.com/nlepage/codyglot/compiler/golang"
	"github.com/spf13/cobra"
)

var _cmd = &cobra.Command{
	Use:   "golang",
	Short: "Start Codyglot Go(lang) compiler",
	RunE: func(_ *cobra.Command, _ []string) error {
		return golang.Compiler.Serve()
	},
}

func init() {
	cmd.Cmd.AddCommand(_cmd)
}
