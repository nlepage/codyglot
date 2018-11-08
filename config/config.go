package config

const (
	// DefaultCleanupBuffer is the default size of cleanup buffer
	DefaultCleanupBuffer = 10
	// DefaultCleanupRoutines is the default number of cleanup routines
	DefaultCleanupRoutines = 2
)

var (
	// CleanupBuffer is the size of the cleanup buffer
	CleanupBuffer int
	// CleanupRoutines is the number of cleanup routines
	CleanupRoutines int
)
