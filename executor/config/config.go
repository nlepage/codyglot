package config

import (
	"github.com/nlepage/codyglot/compiler"
	"github.com/nlepage/codyglot/filestore"
)

const (
	// DefaultPort is the executor default listening port
	DefaultPort = 9090

	DefaultCompilerHost = "localhost" // FIXME correct value
)

var (
	// Port is the executor listening port
	Port int

	Compiler compiler.ClientConfig

	Filestore filestore.ClientConfig
)
