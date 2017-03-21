package mock

import (
	"testing"

	"github.com/sebdah/flipd/source/types"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		key    string
		status int64
	}{
		"enabled:a is enabled": {
			key:    "enabled:a",
			status: types.FeatureStatusEnabled,
		},
		"enabled:b is enabled": {
			key:    "enabled:b",
			status: types.FeatureStatusEnabled,
		},
		"disabled:a is disabled": {
			key:    "disabled:a",
			status: types.FeatureStatusDisabled,
		},
		"disabled:b is disabled": {
			key:    "disabled:b",
			status: types.FeatureStatusDisabled,
		},
	}

	for _, test := range tests {
		store := New()
		feature, err := store.Get(test.key)
		assert.Nil(t, err)
		assert.Equal(t, test.status, feature.Status)
	}
}
