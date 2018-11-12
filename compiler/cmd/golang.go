package cmd

import (
	"github.com/nlepage/codyglot/compiler"
	"github.com/spf13/cobra"
)

func init() {
	addCommand(&cobra.Command{
		Use:   "golang",
		Short: "Start Codyglot Go(lang) compiler",
		RunE: func(_ *cobra.Command, _ []string) error {
			return compiler.Golang(config).Server().Serve()
		},
	}, &config)
}
