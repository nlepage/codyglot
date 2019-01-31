package main

import (
	"github.com/nlepage/codyglot/filestore/cmd"
)

func main() {
	cmd.Filestore.Use = "codyglot-filestore"
	cmd.Filestore.Execute()
}
