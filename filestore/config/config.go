package config

const (
	// DefaultChunkSize is the default chunk size for sending large files
	DefaultChunkSize = 4000000
)

type FileStoreConfig struct {
	// ChunkSize is the chunk size for sending large files
	ChunkSize int
}
