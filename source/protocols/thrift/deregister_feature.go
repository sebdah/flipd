package thrift

import (
	"github.com/agrea/goptr"
	"github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

// DeregisterFeature removes a feature from the flipd service.
//
// Exceptions:
//  - InternalErrorException is thrown for unexpected errors
//  - InvalidInputException is thrown if the feature is not provided
//  - NotFoundException is thrown if the feature could not be found
func (h *Handler) DeregisterFeature(request *flipd.DeregisterFeatureRequest) error {
	if !request.IsSetKey() {
		return &flipd.InvalidInputException{
			Message: goptr.String("required key object not set"),
		}
	}

	err := h.storage.Deregister(string(request.GetKey()))
	if err != nil {
		switch {
		case storage.IsErrNotFound(err):
			return &flipd.NotFoundException{Message: goptr.String(err.Error())}
		default:
			return &flipd.InternalErrorException{Message: goptr.String(err.Error())}
		}
	}

	return nil
}
