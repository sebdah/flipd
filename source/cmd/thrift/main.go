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
	flag.Parse()

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.WithError(err).Panic("Invalid log level set")
	}
	logrus.SetLevel(logLevel)
}

func main() {
	logrus.WithFields(logrus.Fields{"environment": config.Environment}).Info("Starting flipd")

	storage := inmemory.New()

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	address := fmt.Sprintf("%s:%d", config.Address, config.Port)
	secure := false

	if err := runServer(transportFactory, protocolFactory, address, secure, storage); err != nil {
		logrus.WithError(err).Fatal("Could not start the Thrift server")
	}
}
