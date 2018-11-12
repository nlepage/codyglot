package cmd

import (
	"github.com/nlepage/codyglot/executor"
	"github.com/spf13/cobra"
)

func init() {
	addCommand(&cobra.Command{
		Use:   "nodejs",
		Short: "Start Codyglot NodeJS executor",
		RunE: func(_ *cobra.Command, _ []string) error {
			return executor.NodeJS(config).Executor().Serve()
		},
	}, &config)
}
