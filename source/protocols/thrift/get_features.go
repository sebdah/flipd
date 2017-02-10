package thrift

import (
	"github.com/agrea/goptr"
	"github.com/sebdah/flipd/source/storage"
	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

// GetFeatures returns a list of requested features.
//
// Exceptions:
//  - InternalErrorException is thrown for unexpected errors
func (h *Handler) GetFeatures(request *flipd.GetFeaturesRequest) (*flipd.GetFeaturesResponse, error) {
	var err error
	var response = flipd.NewGetFeaturesResponse()

	if request.IsSetKeys() {
		err = h.getFilteredKeys(request.GetKeys(), response)
	} else {
		err = h.getAllKeys(response)
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

// getAllKeys is a helper method fetching all keys from the data store and
// populating the response object.
func (h *Handler) getAllKeys(response *flipd.GetFeaturesResponse) error {
	features, err := h.storage.GetAll()
	if err != nil {
		return &flipd.InternalErrorException{Message: goptr.String(err.Error())}
	}

	for i := range features {
		serializedFeature, err := features[i].SerializeThrift()
		if err != nil {
			return &flipd.InternalErrorException{Message: goptr.String(err.Error())}
		}

		response.Features = append(response.Features, serializedFeature)
	}

	return nil
}

// getFilteredKeys is a helper method fetching selected keys from the data
// store and populating the response object.
func (h *Handler) getFilteredKeys(keys []string, response *flipd.GetFeaturesResponse) error {
	for _, key := range keys {
		feature, err := h.storage.Get(&key)
		if err != nil {
			if !storage.IsErrNotFound(err) {
				return &flipd.InternalErrorException{Message: goptr.String(err.Error())}
			}
		}

		if feature == nil {
			continue
		}

		serializedFeature, err := feature.SerializeThrift()
		if err != nil {
			return &flipd.InternalErrorException{Message: goptr.String(err.Error())}
		}

		response.Features = append(response.Features, serializedFeature)
	}

	return nil
}
