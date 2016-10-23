package types

import "github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"

const (
	// FeatureStatusEnabled indicates that a feature is enabled.
	FeatureStatusEnabled = int64(flipd.FeatureStatus_ENABLED)

	// FeatureStatusDisabled indicates that a feature is disabled.
	FeatureStatusDisabled = int64(flipd.FeatureStatus_DISABLED)
)
