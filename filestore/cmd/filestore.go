package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

// Filestore is the filestore command
var Filestore = &cobra.Command{
	Short: "Codyglot filestore server and client",
	Use:   "filestore",
}

func init() {
	cmd.Cmd.AddCommand(Filestore)
}

func addCommand(_cmd *cobra.Command, config *filestore.Config) {
	_cmd.Flags().IntVar(&config.ChunkSize, "chunk-size", filestore.DefaultChunkSize, "Chunk size for sending large files")
	Filestore.AddCommand(_cmd)
}
