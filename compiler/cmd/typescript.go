package cmd

import (
	"github.com/nlepage/codyglot/compiler"
	"github.com/spf13/cobra"
)

func init() {
	addCommand(&cobra.Command{
		Use:   "typescript",
		Short: "Start Codyglot TypeScript compiler",
		RunE: func(_ *cobra.Command, _ []string) error {
			return compiler.TypeScript(config).Server().Serve()
		},
	}, &config)
}
