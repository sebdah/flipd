package storage

import "github.com/sebdah/flipd/source/types"

// InMemoryBackend stores feature information in memory.
type InMemoryBackend struct {
	// features stored in memory.
	features []types.Feature
}

// NewInMemoryBackend instanciates the memory backend.
func NewInMemoryBackend() *InMemoryBackend { return &InMemoryBackend{} }

// Deregister removes a feature from the storage.
//
// Exceptions:
//  - NotFound error thrown if the key does not exist
func (m *InMemoryBackend) Deregister(key *string) error {
	for i, f := range m.features {
		if *key == f.Key {
			m.features = append(m.features[:i], m.features[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}

// Get a feature from the storage.
//
// Exceptions:
//  - NotFound error thrown if the key does not exist
func (m *InMemoryBackend) Get(key *string) (*types.Feature, error) {
	for _, f := range m.features {
		if f.Key == *key {
			return &f, nil
		}
	}

	return nil, ErrNotFound
}

// Register adds a new feature to the storage.
//
// Exceptions:
//  - Duplicate error is thrown if another feature with the same key already
//    exists.
func (m *InMemoryBackend) Register(feature *types.Feature) error {
	for _, f := range m.features {
		if feature.Key == f.Key {
			return ErrDuplicate
		}
	}

	m.features = append(m.features, *feature)

	return nil
}

// SetStatus updates the status for a given key.
//
// Exceptions:
//  - NotFound error thrown if the key does not exist
func (m *InMemoryBackend) SetStatus(key *string, status *int64) error {
	for i := range m.features {
		if m.features[i].Key == *key {
			m.features[i].Status = *status
			return nil
		}
	}

	return ErrNotFound
}
