package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/spf13/cobra"
)

// Filestore is the filestore command
var Filestore = &cobra.Command{
	Short: "Codyglot filestore server and client",
}

func addCommand(cmd *cobra.Command, config *filestore.Config) {
	cmd.Flags().IntVar(&config.ChunkSize, "chunk-size", filestore.DefaultChunkSize, "Chunk size for sending large files")
	Filestore.AddCommand(cmd)
}
