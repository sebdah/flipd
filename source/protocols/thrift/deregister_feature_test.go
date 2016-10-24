package thrift

import (
	"reflect"
	"testing"

	"github.com/sebdah/flipd/source/storage/inmemory"
	"github.com/sebdah/flipd/source/types"
	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

func TestDeregisterFeature(t *testing.T) {
	features := []*types.Feature{
		&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
		&types.Feature{Key: "some:abc", Status: types.FeatureStatusDisabled},
		&types.Feature{Key: "other:abc", Status: types.FeatureStatusDisabled},
		&types.Feature{Key: "my:key", Status: types.FeatureStatusEnabled},
	}

	tests := []struct {
		key         string
		expectedErr string
		expectedLen int
	}{
		{
			key:         "some:abc",
			expectedErr: "",
			expectedLen: 3,
		},
		{
			key:         "does:not:exist",
			expectedErr: "*flipd.NotFoundException",
			expectedLen: 4,
		},
	}

	for _, test := range tests {
		backend := inmemory.New()
		for _, feature := range features {
			_ = backend.Register(feature)
		}

		service := NewFlipd(backend)

		err := service.DeregisterFeature(&flipd.DeregisterFeatureRequest{Key: flipd.KeyPtr(flipd.Key(test.key))})
		if err != nil {
			if reflect.TypeOf(err).String() != test.expectedErr {
				t.Errorf("expected %s, got %s", test.expectedErr, reflect.TypeOf(err).String())
			}
		}

		response, _ := service.GetFeatures(&flipd.GetFeaturesRequest{})

		if test.expectedLen != len(response.GetFeatures()) {
			t.Errorf("expected %d features in response, got %d", test.expectedLen, len(response.GetFeatures()))
		}
	}
}
