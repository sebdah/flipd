package thrift

import (
	S "github.com/sebdah/flipd/source/storage"
)

// Flipd is implementing the Flipd thrift interfaces.
type Flipd struct {
	storage S.Storager
}

// NewFlipd instanciates a new Flipd thrift service.
func NewFlipd(storage S.Storager) *Flipd {
	return &Flipd{
		storage: storage,
	}
}
