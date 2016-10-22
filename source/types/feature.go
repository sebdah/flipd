package types

// Feature is representing a feature entity in the system.
type Feature struct {
	// Key is a unique identifier for a feature.
	Key string

	// Status indicates the status of the feature, such as enabled or disabled.
	Status int64
}
