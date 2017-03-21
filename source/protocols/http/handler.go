package http

import (
	S "github.com/sebdah/flipd/source/storage"
)

// Handler is implementing the Handler thrift interfaces.
type Handler struct {
	storage S.Storager
}

// NewHandler instanciates a new Handler thrift service.
func NewHandler(storage S.Storager) *Handler {
	return &Handler{
		storage: storage,
	}
}
