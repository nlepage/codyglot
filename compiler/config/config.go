package config

const (
	// DefaultPort is the default listening port of filestore server
	DefaultPort = 9090
)

type CompilerConfig struct {
	// Port is the listening port of filestore server
	Port int
}

var Config CompilerConfig
