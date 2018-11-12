package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/compiler"
	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

var (
	compilerCmd = &cobra.Command{
		Short: "Codyglot compilers",
		Use:   "compiler",
	}
	config compiler.ServerConfig
)

func init() {
	cmd.Cmd.AddCommand(compilerCmd)
}

func addCommand(_cmd *cobra.Command, config *compiler.ServerConfig) {
	_cmd.Flags().IntVarP(&config.Port, "port", "p", compiler.DefaultPort, "Listening port")
	_cmd.Flags().StringVar(&config.Filestore.Host, "filestore-host", filestore.DefaultHost, "Filestore host")
	_cmd.Flags().IntVar(&config.Filestore.Port, "filestore-port", filestore.DefaultPort, "Filestore port")
	_cmd.Flags().IntVar(&config.Filestore.ChunkSize, "filestore-chunk-size", filestore.DefaultChunkSize, "Filestore chunk size")
	compilerCmd.AddCommand(_cmd)
}
