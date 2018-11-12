package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/compiler"
	"github.com/nlepage/codyglot/executor/config"
	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

// Cmd is the executor command group
var Cmd = &cobra.Command{
	Short: "Codyglot executors",
	Use:   "executor",
}

func init() {
	Cmd.PersistentFlags().IntVarP(&config.Port, "port", "p", config.DefaultPort, "Listening port")
	Cmd.PersistentFlags().StringVar(&config.Compiler.Host, "compiler-host", config.DefaultCompilerHost, "Compiler host")
	Cmd.PersistentFlags().IntVar(&config.Compiler.Port, "compiler-port", compiler.DefaultPort, "Compiler port")
	Cmd.PersistentFlags().StringVar(&config.Filestore.Host, "filestore-host", filestore.DefaultHost, "Filestore host")
	Cmd.PersistentFlags().IntVar(&config.Filestore.Port, "filestore-port", filestore.DefaultPort, "Filestore port")
	Cmd.PersistentFlags().IntVar(&config.Filestore.ChunkSize, "filestore-chunk-size", filestore.DefaultChunkSize, "Filestore chunk size")
	cmd.Cmd.AddCommand(Cmd)
}
