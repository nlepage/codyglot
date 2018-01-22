package cmd

import (
	"github.com/Zenika/codyglot/cmd/codyglot"

	// Import subcommands
	_ "github.com/Zenika/codyglot/cmd/router"
	_ "github.com/Zenika/codyglot/cmd/routergateway"
	_ "github.com/Zenika/codyglot/cmd/routerserver"
)

// Cmd is the main command
var Cmd = codyglot.Cmd
