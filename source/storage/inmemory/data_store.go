package inmemory

import (
	"github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/source/types"
)

// DataStore stores feature information in memory.
type DataStore struct {
	// features stored in memory.
	features []types.Feature
}

// New instanciates the memory backend.
func New() *DataStore { return &DataStore{} }

// Deregister removes a feature from the storage.
//
// Exceptions:
//  - NotFound error thrown if the key does not exist
func (d *DataStore) Deregister(key string) error {
	for i, f := range d.features {
		if key == f.Key {
			d.features = append(d.features[:i], d.features[i+1:]...)
			return nil
		}
	}

	return storage.ErrNotFound
}

// Get a feature from the storage.
//
// Exceptions:
//  - NotFound error thrown if the key does not exist
func (d *DataStore) Get(key string) (*types.Feature, error) {
	for _, f := range d.features {
		if f.Key == key {
			return &f, nil
		}
	}

	return nil, storage.ErrNotFound
}

// GetAll returns all features from the storage.
func (d *DataStore) GetAll() ([]*types.Feature, error) {
	var features []*types.Feature

	for i := range d.features {
		features = append(features, &d.features[i])
	}

	return features, nil
}

// Register adds a new feature to the storage.
//
// Exceptions:
//  - Duplicate error is thrown if another feature with the same key already
//    exists.
func (d *DataStore) Register(feature *types.Feature) error {
	for _, f := range d.features {
		if feature.Key == f.Key {
			return storage.ErrDuplicate
		}
	}

	d.features = append(d.features, *feature)

	return nil
}

// SetStatus updates the status for a given key.
//
// Exceptions:
//  - NotFound error thrown if the key does not exist
func (d *DataStore) SetStatus(key string, status int64) error {
	for i := range d.features {
		if d.features[i].Key == key {
			d.features[i].Status = status
			return nil
		}
	}

	return storage.ErrNotFound
}
