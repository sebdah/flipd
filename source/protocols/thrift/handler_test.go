package thrift

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	assert.IsType(t, &Handler{}, NewHandler(nil))
}
