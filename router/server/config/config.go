package config

import "time"

const (
	// DefaultPort is the default router listening port
	DefaultPort = 9090
)

var (
	// DefaultExecutors is the list of default executors endpoints
	DefaultExecutors = []string{"golang", "nodejs"}

	// Port is the router listening port
	Port int
	// Executors is the list of executors endpoints
	Executors []string
	// PingInterval is the time between executor pings (in ms)
	PingInterval = 1000 * time.Millisecond
	// MaxBackoff is the maximum backoff time for executor ping (in ms)
	MaxBackoff = 51200 * time.Millisecond
)
