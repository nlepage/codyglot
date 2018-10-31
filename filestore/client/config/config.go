package config

import (
	"github.com/nlepage/codyglot/filestore/config"
)

const (
	// DefaultHost is the default file store server host
	DefaultHost = "localhost"

	// DefaultPort is the default file store server port
	DefaultPort = 9090
)

type FileStoreClientConfig struct {
	config.FileStoreConfig

	// Host is the file store server host
	Host string

	// Port is the file store server port
	Port int
}
