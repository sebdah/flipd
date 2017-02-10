package main

import (
	"flag"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/sebdah/flipd/source/config"
	"github.com/sebdah/flipd/source/storage/inmemory"
)

func init() {
	flag.StringVar(&config.Environment, "environment", "development", "Application environment name")
	flag.StringVar(&config.Host, "host", "localhost", "Host name the service should bind to")
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
	logrus.WithFields(logrus.Fields{
		"environment": config.Environment,
		"port":        config.Port,
	}).Info("Starting flipd")

	storage := inmemory.New()

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	secure := false

	if err := runServer(transportFactory, protocolFactory, address, secure, storage); err != nil {
		logrus.WithError(err).Fatal("Could not start the Thrift server")
	}
}
