package config

const (
	// DefaultPort is the default router gateway listening port
	DefaultPort = 8080
	// DefaultEndpoint is the default router gateway endpoint
	DefaultEndpoint = "localhost:9090"
)

var (
	// Port is the router gateway listening port
	Port int
	// Endpoint is the router gateway endpoint
	Endpoint string
)
