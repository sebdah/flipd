package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/sebdah/flipd/source/config"
)

func init() {
	flag.StringVar(&config.Environment, "environment", "development", "Application environment name")
	flag.StringVar(&config.Address, "address", "localhost", "Address the service should bind to")
	flag.IntVar(&config.Port, "port", 9090, "Port number the service should bind to")
	flag.StringVar(&config.LogLevel, "log-level", "debug", "Log level used in the application")
	flag.Parse()

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.WithError(err).Panic("Invalid log level set")
	}
	logrus.SetLevel(logLevel)
}

func main() {

}
