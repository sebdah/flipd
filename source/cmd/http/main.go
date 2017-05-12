package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-zoo/bone"
	"github.com/sebdah/flipd/source/config"
	httpapi "github.com/sebdah/flipd/source/protocols/http"
	"github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/source/storage/inmemory"
	"github.com/sebdah/flipd/source/storage/mock"
)

var storageType string

func init() {
	flag.StringVar(&config.Environment, "environment", "development", "Application environment name")
	flag.StringVar(&config.Host, "host", "localhost", "Host the service should bind to")
	flag.IntVar(&config.Port, "port", 9090, "Port number the service should bind to")
	flag.StringVar(&config.LogLevel, "log-level", "debug", "Log level used in the application")
	flag.StringVar(&storageType, "storage", "inmemory", "Storage backend to use. One of inmemory or mock.")
	flag.Parse()

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.WithError(err).Panic("Invalid log level set")
	}
	logrus.SetLevel(logLevel)
}

func main() {
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	logrus.WithFields(logrus.Fields{
		"environment": config.Environment,
		"address":     address,
	}).Info("Starting flipd http service")

	var storageBackend storage.Storager

	switch storageType {
	case "mock":
		storageBackend = mock.New()
	case "inmemory":
		storageBackend = inmemory.New()
	}

	api := httpapi.NewHandler(storageBackend)

	mux := bone.New()
	assert.Nil(t, err)
	mux.Get("/features", http.HandlerFunc(api.GetFeatures))

	err := http.ListenAndServe(address, mux)
	if err != nil {
		logrus.WithError(err).Fatal("Could not start http server")
	}
}
