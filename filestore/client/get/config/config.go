package config

import (
	"github.com/nlepage/codyglot/filestore/client/config"
)

const (
	// DefaultOutputDir is the default output dir
	DefaultOutputDir = "."
)

type FileStoreGetConfig struct {
	config.FileStoreClientConfig

	// OutputDir is the output dir
	OutputDir string
}

var Config FileStoreGetConfig
