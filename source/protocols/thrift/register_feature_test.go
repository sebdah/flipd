package thrift

import (
	"reflect"
	"testing"

	"github.com/sebdah/flipd/source/storage/inmemory"
	"github.com/sebdah/flipd/thrift/generated-code/gen-go/flipd"
)

func TestRegisterFeature(t *testing.T) {
	service := NewHandler(inmemory.New())

	err := service.RegisterFeature(
		&flipd.RegisterFeatureRequest{
			Feature: &flipd.Feature{
				Key:    flipd.KeyPtr(flipd.Key("my:key")),
				Status: flipd.FeatureStatus_ENABLED,
			},
		})
	if err != nil {
		t.Errorf("expected nil, got %s", err.Error())
	}

	err = service.RegisterFeature(
		&flipd.RegisterFeatureRequest{
			Feature: &flipd.Feature{
				Key:    flipd.KeyPtr(flipd.Key("my:key")),
				Status: flipd.FeatureStatus_ENABLED,
			},
		})
	if reflect.TypeOf(err).String() != "*flipd.DuplicateException" {
		t.Errorf("expected *flipd.DuplicateException, got %s", reflect.TypeOf(err).String())
	}

	err = service.RegisterFeature(&flipd.RegisterFeatureRequest{})
	if reflect.TypeOf(err).String() != "*flipd.InvalidInputException" {
		t.Errorf("expected *flipd.InvalidInputException, got %s", reflect.TypeOf(err).String())
	}

	err = service.RegisterFeature(
		&flipd.RegisterFeatureRequest{
			Feature: &flipd.Feature{
				Key: flipd.KeyPtr(flipd.Key("other:key")),
			},
		})
	if err != nil {
		t.Errorf("expected nil, got %s", err.Error())
	}
}
