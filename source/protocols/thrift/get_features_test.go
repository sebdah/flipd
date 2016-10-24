package thrift

import (
	"testing"

	"github.com/sebdah/flipd/source/storage/inmemory"
	"github.com/sebdah/flipd/source/types"
	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

func TestGetFeatures(t *testing.T) {
	features := []*types.Feature{
		&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
		&types.Feature{Key: "some:abc", Status: types.FeatureStatusDisabled},
		&types.Feature{Key: "other:abc", Status: types.FeatureStatusDisabled},
		&types.Feature{Key: "my:key", Status: types.FeatureStatusEnabled},
	}

	backend := inmemory.New()
	for _, feature := range features {
		_ = backend.Register(feature)
	}

	service := NewFlipd(backend)

	tests := []struct {
		keys        []string
		expectedErr error
		expectedLen int
	}{
		{
			keys:        []string{},
			expectedErr: nil,
			expectedLen: 0,
		},
		{
			keys:        nil,
			expectedErr: nil,
			expectedLen: 4,
		},
		{
			keys:        []string{"some:key"},
			expectedErr: nil,
			expectedLen: 1,
		},
		{
			keys:        []string{"unknown:key"},
			expectedErr: nil,
			expectedLen: 0,
		},
		{
			keys:        []string{"some:key", "my:key", "does:not:exist"},
			expectedErr: nil,
			expectedLen: 2,
		},
	}

	for _, test := range tests {
		response, err := service.GetFeatures(&flipd.GetFeaturesRequest{Keys: test.keys})
		if test.expectedErr == nil && err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}
		if test.expectedErr != nil && err == nil {
			t.Errorf("expected %s, got nil", test.expectedErr.Error())
		}

		if test.expectedErr != nil && err != nil && test.expectedErr.Error() != err.Error() {
			t.Errorf("expected %s, got %s", test.expectedErr.Error(), err.Error())
		}

		if test.expectedLen != len(response.GetFeatures()) {
			t.Errorf("expected %d features in response, got %d", test.expectedLen, len(response.GetFeatures()))
		}
	}
}
