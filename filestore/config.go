package filestore

const (
	// DefaultChunkSize is the default chunk size for sending large files
	DefaultChunkSize = 4000000

	// DefaultHost is the default file store server host
	DefaultHost = "filestore"

	// DefaultPort is the default file store server port
	DefaultPort = 9090

	// DefaultOutputDir is the default output dir
	DefaultOutputDir = "."

	// DefaultRoot is the default root directory of filestore
	DefaultRoot = "/var/codyglot/filestore"
)

type Config struct {
	// ChunkSize is the chunk size for sending large files
	ChunkSize int
}

type ClientConfig struct {
	Config

	// Host is the file store server host
	Host string

	// Port is the file store server port
	Port int
}

type GetConfig struct {
	ClientConfig

	// OutputDir is the output dir
	OutputDir string
}

type ServerConfig struct {
	Config

	// Port is the listening port of filestore server
	Port int

	// Root is the root directory of filestore
	Root string
}
