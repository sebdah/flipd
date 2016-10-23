package types

import (
	"errors"
	"testing"

	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

func TestFeatureDeserialize(t *testing.T) {
	tests := []struct {
		input           *flipd.Feature
		expectedFeature *Feature
		expectedErr     error
	}{
		{
			input:           &flipd.Feature{Key: flipd.KeyPtr(flipd.Key("my:key")), Status: flipd.FeatureStatus_DISABLED},
			expectedFeature: &Feature{Key: "my:key", Status: FeatureStatusDisabled},
			expectedErr:     nil,
		},
		{
			input:           &flipd.Feature{Key: flipd.KeyPtr(flipd.Key("my:key")), Status: flipd.FeatureStatus_ENABLED},
			expectedFeature: &Feature{Key: "my:key", Status: FeatureStatusEnabled},
			expectedErr:     nil,
		},
		{
			input:           nil,
			expectedFeature: nil,
			expectedErr:     errors.New("cannot deserialize nil"),
		},
		{
			input:           &flipd.Feature{Key: nil, Status: flipd.FeatureStatus_DISABLED},
			expectedFeature: nil,
			expectedErr:     errors.New("missing required key attribute"),
		},
	}

	for _, test := range tests {
		feature := NewFeature()
		err := feature.DeserializeThrift(test.input)
		switch {
		case err != nil && test.expectedErr != nil:
			if err.Error() != test.expectedErr.Error() {
				t.Errorf("expected %v, got %v", test.expectedErr, err)
			}
		case err != test.expectedErr:
			t.Errorf("expected %v, got %v", test.expectedErr, err)
		}

		if test.expectedFeature != nil {
			if test.expectedFeature.Key != string(test.input.GetKey()) {
				t.Errorf("expected %s, got %s", test.expectedFeature.Key, string(test.input.GetKey()))
			}

			if test.expectedFeature.Status != int64(test.input.GetStatus()) {
				t.Errorf("expected %d, got %d", test.expectedFeature.Status, int64(test.input.GetStatus()))
			}
		}
	}
}

func TestFeatureSerialize(t *testing.T) {
	tests := []struct {
		feature        *Feature
		expectedResult *flipd.Feature
		expectedErr    error
	}{
		{
			feature:        &Feature{Key: "my:key", Status: FeatureStatusEnabled},
			expectedResult: &flipd.Feature{Key: flipd.KeyPtr(flipd.Key("my:key")), Status: flipd.FeatureStatus_ENABLED},
			expectedErr:    nil,
		},
		{
			feature:        &Feature{Key: "my:key"},
			expectedResult: &flipd.Feature{Key: flipd.KeyPtr(flipd.Key("my:key")), Status: flipd.FeatureStatus_DISABLED},
			expectedErr:    nil,
		},
		{
			feature:        &Feature{},
			expectedResult: nil,
			expectedErr:    errors.New("missing required key attribute"),
		},
	}

	for _, test := range tests {
		result, err := test.feature.SerializeThrift()
		if test.expectedErr == nil && err != nil {
			t.Errorf("expected err, got %s", err.Error())
		}

		if test.expectedResult == nil && result != nil {
			t.Errorf("expected nil, got %v", result)
		}

		if test.expectedResult != nil {
			if test.expectedResult.GetKey() != result.GetKey() {
				t.Errorf("expected %s, got %s", string(test.expectedResult.GetKey()), result.Key)
			}

			if test.feature.Status != 0 && test.expectedResult.GetStatus() != result.GetStatus() {
				t.Errorf("expected %v, got %v", test.expectedResult.GetStatus(), result.GetStatus())
			}

			if test.feature.Status == 0 && int64(result.GetStatus()) != FeatureStatusDisabled {
				t.Errorf("expected %v, got %v", test.expectedResult.GetStatus(), result.GetStatus())
			}
		}
	}
}
