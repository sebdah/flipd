package config

var (
	// Environment holds the current environment name. E.g. "development",
	// "test" or "production".
	Environment string

	// Host the service should bind to.
	Host string

	// Port is the port number the service should bind to.
	Port int

	// LogLevel used in the application.
	// Valid levels are:
	//  - debug
	//  - info
	//  - warning
	//  - error
	//  - fatal
	//  - panic
	LogLevel string
)
