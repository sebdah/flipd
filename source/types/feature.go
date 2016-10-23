package types

import (
	"errors"

	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

// Feature is representing a feature entity in the system.
type Feature struct {
	// Key is a unique identifier for a feature.
	Key string

	// Status indicates the status of the feature, such as enabled or disabled.
	Status int64
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
