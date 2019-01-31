package filestore

import (
	"github.com/nlepage/codyglot/cmd"
	fscmd "github.com/nlepage/codyglot/filestore/cmd"
)

func init() {
	fscmd.Filestore.Use = "filestore"
	cmd.Cmd.AddCommand(fscmd.Filestore)
}
