package main

import (
	"crypto/tls"

	"github.com/Sirupsen/logrus"
	"github.com/apache/thrift/lib/go/thrift"
	T "github.com/sebdah/flipd/source/protocols/thrift"
	S "github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, storage S.Storager) error {
	var transport thrift.TServerTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		if cert, e := tls.LoadX509KeyPair("server.crt", "server.key"); e == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			return e
		}

		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		return err
	}

	processor := flipd.NewFlipdProcessor(T.NewHandler(storage))
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	logrus.WithFields(logrus.Fields{"address": addr}).Info("Starting Thrift server")

	return server.Serve()
}
