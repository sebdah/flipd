package types

import (
	"testing"

	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

func TestFeatureStatus(t *testing.T) {
	if FeatureStatusDisabled != int64(flipd.FeatureStatus_DISABLED) {
		t.Errorf("expected %d, got %d", FeatureStatusDisabled, int64(flipd.FeatureStatus_DISABLED))
	}

	if FeatureStatusEnabled != int64(flipd.FeatureStatus_ENABLED) {
		t.Errorf("expected %d, got %d", FeatureStatusEnabled, int64(flipd.FeatureStatus_ENABLED))
	}
}
