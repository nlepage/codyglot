package cmd

import (
	"github.com/nlepage/codyglot/executor"
	"github.com/spf13/cobra"
)

func init() {
	addCommand(&cobra.Command{
		Use:   "golang",
		Short: "Start Codyglot Go(lang) executor",
		RunE: func(_ *cobra.Command, _ []string) error {
			return executor.Golang(config).Executor().Serve()
		},
	}, &config)
}
