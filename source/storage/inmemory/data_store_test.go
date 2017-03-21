package inmemory

import (
	"testing"

	"github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/source/types"
)

func TestDeregisterFeature(t *testing.T) {
	tests := []struct {
		features    []*types.Feature
		key         string
		expectedErr error
	}{
		{
			features: []*types.Feature{
				&types.Feature{Key: "a:key", Status: types.FeatureStatusEnabled},
				&types.Feature{Key: "a:abc", Status: types.FeatureStatusDisabled},
			},
			key:         "a:key",
			expectedErr: nil,
		},
		{
			features: []*types.Feature{
				&types.Feature{Key: "a:key", Status: types.FeatureStatusEnabled},
				&types.Feature{Key: "a:abc", Status: types.FeatureStatusDisabled},
				&types.Feature{Key: "b:abd", Status: types.FeatureStatusDisabled},
			},
			key:         "new:key",
			expectedErr: storage.ErrNotFound,
		},
	}

	for _, test := range tests {
		dataStore := New()

		for _, feature := range test.features {
			dataStore.Register(feature)
		}

		err := dataStore.Deregister(test.key)
		if err != test.expectedErr {
			switch {
			case test.expectedErr == nil:
				t.Errorf("expected nil, got %s", err.Error())
			case err == nil:
				t.Errorf("expected %s, got nil", test.expectedErr.Error())
			}
		}

		_, err = dataStore.Get(test.key)
		if !storage.IsErrNotFound(err) {
			t.Errorf("expected %s, got %s", storage.ErrNotFound.Error(), err.Error())
		}
	}
}

func TestGetFeature(t *testing.T) {
	tests := []struct {
		features    []*types.Feature
		key         string
		expectedErr error
	}{
		{
			features: []*types.Feature{
				&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
				&types.Feature{Key: "some:abc", Status: types.FeatureStatusDisabled},
			},
			key:         "some:key",
			expectedErr: nil,
		},
		{
			features: []*types.Feature{
				&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
				&types.Feature{Key: "some:abc", Status: types.FeatureStatusDisabled},
			},
			key:         "other:key",
			expectedErr: storage.ErrNotFound,
		},
	}

	for _, test := range tests {
		dataStore := New()

		for _, feature := range test.features {
			dataStore.Register(feature)
		}

		feature, err := dataStore.Get(test.key)
		if err != test.expectedErr {
			switch {
			case test.expectedErr == nil:
				t.Errorf("expected nil, got %s", err.Error())
			case err == nil:
				t.Errorf("expected %s, got nil", test.expectedErr.Error())
			}
		}

		if feature != nil && feature.Key != test.key {
			t.Errorf("expected %s, got %s", feature.Key, test.key)
		}
	}
}

func TestGetAllFeature(t *testing.T) {
	tests := []struct {
		features []*types.Feature
	}{
		{
			features: []*types.Feature{
				&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
				&types.Feature{Key: "some:abc", Status: types.FeatureStatusDisabled},
				&types.Feature{Key: "other:abc", Status: types.FeatureStatusDisabled},
				&types.Feature{Key: "my:key", Status: types.FeatureStatusEnabled},
			},
		},
		{
			features: []*types.Feature{},
		},
	}

	for _, test := range tests {
		dataStore := New()

		for _, feature := range test.features {
			dataStore.Register(feature)
		}

		result, err := dataStore.GetAll()
		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}

		if len(result) != len(test.features) {
			t.Errorf("expected %d items in the result, got %d", len(test.features), len(result))
		}
	}
}

func TestRegisterFeature(t *testing.T) {
	feature1 := &types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled}
	feature2 := &types.Feature{Key: "some:abc", Status: types.FeatureStatusEnabled}

	dataStore := New()

	err := dataStore.Register(feature1)
	if err != nil {
		t.Errorf("expected nil, got %s", err.Error())
	}

	err = dataStore.Register(feature1)
	if !storage.IsErrDuplicate(err) {
		switch {
		case err == nil:
			t.Error("expected duplication error, got nil")
		default:
			t.Errorf("expected duplication error, got %s", err.Error())
		}
	}

	err = dataStore.Register(feature2)
	if err != nil {
		t.Errorf("expected nil, got %s", err.Error())
	}
}

func TestSetStatus(t *testing.T) {
	tests := []struct {
		features    []*types.Feature
		key         string
		status      int64
		expectedErr error
	}{
		{
			features: []*types.Feature{
				&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
				&types.Feature{Key: "some:abc", Status: types.FeatureStatusDisabled},
			},
			key:         "some:key",
			status:      types.FeatureStatusDisabled,
			expectedErr: nil,
		},
		{
			features: []*types.Feature{
				&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
				&types.Feature{Key: "some:abc", Status: types.FeatureStatusDisabled},
			},
			key:         "other:key",
			status:      types.FeatureStatusDisabled,
			expectedErr: storage.ErrNotFound,
		},
	}

	for _, test := range tests {
		dataStore := New()

		for _, feature := range test.features {
			dataStore.Register(feature)
		}

		err := dataStore.SetStatus(test.key, test.status)
		if err != test.expectedErr {
			switch {
			case test.expectedErr == nil:
				t.Errorf("expected nil, got %s", err.Error())
			case err == nil:
				t.Errorf("expected %s, got nil", test.expectedErr.Error())
			}
		}

		if test.expectedErr == nil {
			feature, err := dataStore.Get(test.key)
			if err != nil {
				t.Errorf("expected nil, got %s", err.Error())
			}

			if feature.Status != test.status {
				t.Errorf("expected %d, got %d", test.status, feature.Status)
			}
		}
	}
}
