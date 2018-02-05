package config

const (
	// DefaultPort is the default router listening port
	DefaultPort = 9090
)

var (
	// DefaultExecutors is the list of default executors endpoints
	DefaultExecutors = []string{"golang:9090", "nodejs:9090"}

	// Port is the router listening port
	Port int
	// Executors is the list of executors endpoints
	Executors []string
)
