package mock

import (
	"github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/source/types"
)

// New returns a mock store with some prepopulated keys.
func New() storage.Storager {
	return &DataStore{}
}

// DataStore is the mock data store object.
type DataStore struct {
	DeregisterErr error
	GetAllErr     error
	GetAllFeature []*types.Feature
	GetErr        error
	GetFeature    *types.Feature
	RegisterErr   error
	SetStatusErr  error
}

// Deregister mocks the Deregister method.
func (d *DataStore) Deregister(name string) error {
	return d.DeregisterErr
}

// Get mocks the get method.
func (d *DataStore) Get(string) (*types.Feature, error) {
	return d.GetFeature, g.GetErr
}

// GetAll mocks the GetAll method.
func (d *DataStore) GetAll() ([]*types.Feature, error) {
	return d.GetAllFeature, d.GetAllErr
}

// Register mocks the Register method.
func (d *DataStore) Register(feature *types.Feature) error {
	return d.RegisterErr
}

// SetStatus mocks the SetStatus method.
func (d *DataStore) SetStatus(name string, status int64) error {
	return d.RegisterErr
}
