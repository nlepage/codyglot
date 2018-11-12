package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/compiler"
	"github.com/nlepage/codyglot/executor"
	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

var (
	executorCmd = &cobra.Command{
		Short: "Codyglot executors",
		Use:   "executor",
	}
	config executor.Config
)

func init() {
	cmd.Cmd.AddCommand(executorCmd)
}

func addCommand(_cmd *cobra.Command, config *executor.Config) {
	_cmd.Flags().IntVarP(&config.Port, "port", "p", executor.DefaultPort, "Listening port")
	_cmd.Flags().StringVar(&config.Compiler.Host, "compiler-host", executor.DefaultCompilerHost, "Compiler host")
	_cmd.Flags().IntVar(&config.Compiler.Port, "compiler-port", compiler.DefaultPort, "Compiler port")
	_cmd.Flags().StringVar(&config.Filestore.Host, "filestore-host", filestore.DefaultHost, "Filestore host")
	_cmd.Flags().IntVar(&config.Filestore.Port, "filestore-port", filestore.DefaultPort, "Filestore port")
	_cmd.Flags().IntVar(&config.Filestore.ChunkSize, "filestore-chunk-size", filestore.DefaultChunkSize, "Filestore chunk size")
	executorCmd.AddCommand(_cmd)
}
