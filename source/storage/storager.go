package storage

import "github.com/sebdah/flipd/source/types"

// Storager defines the methods available on the storage backends.
type Storager interface {
	// Get returns a feature given it's key.
	Get(*string) (*types.Feature, error)

	// Deregister removes a key from the data store.
	Deregister(*string) error

	// Register adds a feature to the data store.
	Register(*types.Feature) error

	// Set status updates the status of a feature in the data store.
	SetStatus(*string, *int64) error
}
