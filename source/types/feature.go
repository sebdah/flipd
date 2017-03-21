package types

import (
	"errors"

	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

// Feature is representing a feature entity in the system.
type Feature struct {
	// Key is a unique identifier for a feature.
	Key string `json:"key"`

	// Status indicates the status of the feature, such as enabled or disabled.
	Status int64 `json:"status"`
}

// NewFeature returns a new instance of a feature.
func NewFeature() *Feature { return &Feature{} }

// DeserializeThrift converts a Thrift feature object to a Feature.
func (f *Feature) DeserializeThrift(feature *flipd.Feature) error {
	if feature == nil {
		return errors.New("cannot deserialize nil")
	}

	if !feature.IsSetKey() {
		return errors.New("missing required key attribute")
	}

	f.Key = string(feature.GetKey())
	f.Status = int64(feature.GetStatus())

	return nil
}

// SerializeThrift converts a feature to a Thrift feature object.
func (f *Feature) SerializeThrift() (*flipd.Feature, error) {
	if f.Key == "" {
		return nil, errors.New("missing required key attribute")
	}

	feature := flipd.NewFeature()
	feature.Key = flipd.KeyPtr(flipd.Key(f.Key))

	if f.Status != 0 {
		feature.Status = flipd.FeatureStatus(f.Status)
	}

	return feature, nil
}
