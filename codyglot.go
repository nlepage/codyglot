package main

import (
	"github.com/nlepage/codyglot/cmd"

	// Import sub commands
	_ "github.com/nlepage/codyglot/compiler/cmd"
	_ "github.com/nlepage/codyglot/executor/cmd"
	_ "github.com/nlepage/codyglot/filestore/cmd"
	_ "github.com/nlepage/codyglot/graphql/cmd"
	_ "github.com/nlepage/codyglot/router/gateway/cmd"
	_ "github.com/nlepage/codyglot/router/server/cmd"
)

func main() {
	cmd.Cmd.Execute()
}
