package http_client_test

import (
	"home_sensor_lambda/internal/http_client"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	MockGet func(url string) (*http.Response, error)
}

func (c *MockClient) Get(url string) (*http.Response, error) {
	return c.MockGet(url)
}

func TestMakeRequest(t *testing.T) {
	mockClient := &MockClient{
		MockGet: func(url string) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`OK`)),
			}, nil
		},
	}

	body, err := http_client.Get("https://example.com", mockClient)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if body != "OK" {
		t.Fatalf("expected OK, got %s", body)
	}
}
