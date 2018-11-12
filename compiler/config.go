package compiler

import (
	"github.com/nlepage/codyglot/filestore"
)

const (
	// DefaultPort is the default listening port of filestore server
	DefaultPort = 9090
)

type ServerConfig struct {
	// Port is the listening port of filestore server
	Port int

	Filestore filestore.ClientConfig
}

type ClientConfig struct {
	Host string
	Port int
}
