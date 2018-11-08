package cmd

import (
	"github.com/nlepage/codyglot/compiler"
	"github.com/spf13/cobra"
)

var (
	config compiler.Config

	golangCmd = &cobra.Command{
		Use:   "golang",
		Short: "Start Codyglot Go(lang) compiler",
		RunE: func(_ *cobra.Command, _ []string) error {
			return compiler.Golang(config).Compiler().Serve()
		},
	}
)

func init() {
	addCommand(golangCmd, &config)
}
