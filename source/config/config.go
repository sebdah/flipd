package config

import (
	"flag"
	"os"
)

var (
	// Environment holds the current environment name. E.g. "development",
	// "test" or "production".
	Environment string

	// Address the service should bind to.
	Address string

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

func init() {
	flag.StringVar(&Environment, "environment", os.Getenv("APP_ENV"), "Application environment name")
	flag.StringVar(&Address, "address", "localhost", "Address the service should bind to")
	flag.IntVar(&Port, "port", 9090, "Port number the service should bind to")
	flag.StringVar(&LogLevel, "log-level", "debug", "Log level used in the application")

	if Environment == "" {
		Environment = "development"
	}
}
