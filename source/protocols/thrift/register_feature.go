package thrift

import (
	"github.com/agrea/goptr"
	"github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/source/types"
	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

// RegisterFeature adds a new feature to the flipd service. If the feature
// already exists, a DuplicationException will be thrown.
//
// Exceptions:
//  - InternalErrorException is thrown for unexpected errors
//  - InvalidInputException is thrown if the feature is not provided
//  - DuplicateException is thrown if the feature is already added
func (h *Handler) RegisterFeature(request *flipd.RegisterFeatureRequest) error {
	if request.Feature == nil {
		return &flipd.InvalidInputException{
			Message: goptr.String("required feature object not set"),
		}
	}

	feature := types.NewFeature()
	feature.DeserializeThrift(request.Feature)

	err := h.storage.Register(feature)
	if err != nil {
		switch {
		case storage.IsErrDuplicate(err):
			return &flipd.DuplicateException{Message: goptr.String(err.Error())}
		default:
			return &flipd.InternalErrorException{Message: goptr.String(err.Error())}
		}
	}

	return nil
}
