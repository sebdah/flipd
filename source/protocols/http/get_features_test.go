package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sebdah/flipd/source/storage/inmemory"
	"github.com/sebdah/flipd/source/types"
	"github.com/stretchr/testify/assert"
)

func TestGetFeatures(t *testing.T) {
	tests := []struct {
		features []*types.Feature
	}{
		{
			features: []*types.Feature{},
		},
		{
			features: []*types.Feature{
				&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
			},
		},
		{
			features: []*types.Feature{
				&types.Feature{Key: "some:key", Status: types.FeatureStatusDisabled},
			},
		},
		{
			features: []*types.Feature{
				&types.Feature{Key: "some:key", Status: types.FeatureStatusEnabled},
				&types.Feature{Key: "some:abc", Status: types.FeatureStatusDisabled},
				&types.Feature{Key: "other:abc", Status: types.FeatureStatusDisabled},
				&types.Feature{Key: "my:key", Status: types.FeatureStatusEnabled},
			},
		},
	}

	for _, test := range tests {
		backend := inmemory.New()
		for _, feature := range test.features {
			err := backend.Register(feature)
			assert.Nil(t, err)
		}

		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/features", nil)
		assert.Nil(t, err)

		h := http.HandlerFunc(NewHandler(backend).GetFeatures)
		h.ServeHTTP(recorder, req)

		response := &getFeaturesResponse{}
		json.NewDecoder(recorder.Body).Decode(response)

		assert.Equal(t, len(test.features), len(response.Features))

		for i := range test.features {
			assert.Equal(t, test.features[i], response.Features[i])
		}
	}
}
