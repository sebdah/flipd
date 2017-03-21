package mock

import (
	"github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/source/storage/inmemory"
	"github.com/sebdah/flipd/source/types"
)

// New returns a mock store with some prepopulated keys.
func New() storage.Storager {
	store := inmemory.New()

	store.Register(&types.Feature{
		Key:    "enabled:a",
		Status: types.FeatureStatusEnabled,
	})
	store.Register(&types.Feature{
		Key:    "enabled:b",
		Status: types.FeatureStatusEnabled,
	})
	store.Register(&types.Feature{
		Key:    "disabled:a",
		Status: types.FeatureStatusDisabled,
	})
	store.Register(&types.Feature{
		Key:    "disabled:b",
		Status: types.FeatureStatusDisabled,
	})

	return store
}
