package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/compiler"
	"github.com/nlepage/codyglot/filestore"
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

// FIXME bind filestore client config
func addCommand(_cmd *cobra.Command, config *compiler.ServerConfig) {
	Compiler.Flags().IntVarP(&config.Port, "port", "p", compiler.DefaultPort, "Listening port")
	Compiler.Flags().StringVar(&config.Filestore.Host, "filestore-host", filestore.DefaultHost, "Filestore host")
	Compiler.Flags().IntVar(&config.Filestore.Port, "filestore-port", filestore.DefaultPort, "Filestore port")
	Compiler.Flags().IntVar(&config.Filestore.ChunkSize, "filestore-chunk-size", filestore.DefaultChunkSize, "Filestore chunk size")
	Compiler.AddCommand(_cmd)
}
