package cmd

import (
	"github.com/nlepage/codyglot/cmd/codyglot"

	// Import subcommands
	_ "github.com/nlepage/codyglot/cmd/executor"
	_ "github.com/nlepage/codyglot/cmd/executor/golang"
	_ "github.com/nlepage/codyglot/cmd/router"
	_ "github.com/nlepage/codyglot/cmd/router/gateway"
	_ "github.com/nlepage/codyglot/cmd/router/server"
)

// Cmd is the main command
var Cmd = codyglot.Cmd
