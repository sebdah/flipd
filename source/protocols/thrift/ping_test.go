package thrift

import "testing"

func TestPing(t *testing.T) {
	service := NewHandler(nil)
	response, err := service.Ping()

	if err != nil {
		t.Errorf("expected nil, got %s", err.Error())
	}

	if response != "pong" {
		t.Errorf("expected pong, got %s", response)
	}
}
