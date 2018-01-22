package cmd

import (
	"github.com/Zenika/codyglot/cmd/codyglot"

	// Import subcommands
	_ "github.com/Zenika/codyglot/cmd/executor"
	_ "github.com/Zenika/codyglot/cmd/executor/golang"
	_ "github.com/Zenika/codyglot/cmd/router"
	_ "github.com/Zenika/codyglot/cmd/router/gateway"
	_ "github.com/Zenika/codyglot/cmd/router/server"
)

// Cmd is the main command
var Cmd = codyglot.Cmd
