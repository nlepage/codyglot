package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/compiler"
	"github.com/spf13/cobra"
)

// Compiler is the compiler command group
var Compiler = &cobra.Command{
	Short: "Codyglot compilers",
	Use:   "compiler",
}

func init() {
	cmd.Cmd.AddCommand(Compiler)
}

func addCommand(_cmd *cobra.Command, config *compiler.Config) {
	Compiler.Flags().IntVarP(&config.Port, "port", "p", compiler.DefaultPort, "Listening port")
	Compiler.AddCommand(_cmd)
}
