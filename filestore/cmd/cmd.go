package cmd

import (
	"github.com/nlepage/codyglot/cmd"
	"github.com/nlepage/codyglot/filestore/config"
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

func AddCommand(_cmd *cobra.Command, cfg *config.FileStoreConfig) {
	_cmd.Flags().IntVar(&cfg.ChunkSize, "chunk-size", config.DefaultChunkSize, "Chunk size for sending large files")
	Cmd.AddCommand(_cmd)
}
