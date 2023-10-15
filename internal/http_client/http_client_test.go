package http_client_test

import (
	"home_sensor_lambda/internal/http_client"
	"home_sensor_lambda/test/mocks"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	mockClient := mocks.HttpGetSuccess()
	body, err := http_client.Get("https://example.com", mockClient)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if body != "OK" {
		t.Fatalf("expected OK, got %s", body)
	}
}
