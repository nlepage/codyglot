package cmd

import (
	"github.com/nlepage/codyglot/filestore"
	"github.com/nlepage/codyglot/cmd"
	"github.com/spf13/cobra"
)

// Cmd is the filestore command group
var Cmd = &cobra.Command{
	Short: "Codyglot filestore commands",
	Use:   "filestore",
}

func init() {
	cmd.Cmd.AddCommand(Cmd)
}

func AddCommand(_cmd *cobra.Command, config *filestore.Config) {
	_cmd.Flags().IntVar(&config.ChunkSize, "chunk-size", filestore.DefaultChunkSize, "Chunk size for sending large files")
	Cmd.AddCommand(_cmd)
}
