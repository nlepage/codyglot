package config

import (
	"github.com/nlepage/codyglot/filestore/config"
)

const (
	// DefaultPort is the default listening port of filestore server
	DefaultPort = 9090

	// DefaultRoot is the default root directory of filestore
	DefaultRoot = "/var/codyglot/filestore"
)

type FileStoreServerConfig struct {
	config.FileStoreConfig

	// Port is the listening port of filestore server
	Port int

	// Root is the root directory of filestore
	Root string
}

var Config FileStoreServerConfig
