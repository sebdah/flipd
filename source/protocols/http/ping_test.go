package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-zoo/bone"
	"github.com/sebdah/flipd/source/storage/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	api := NewHandler(inmemory.New())
	mux := bone.New()
	mux.Get("/ping", http.HandlerFunc(api.Ping))

	server := httptest.NewServer(mux)
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/ping", server.URL))
	assert.Nil(t, err)
	defer res.Body.Close()

	response := &pingResponse{}
	err = json.NewDecoder(res.Body).Decode(response)
	assert.Nil(t, err)

	assert.Equal(t, "Pong", response.Message)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
