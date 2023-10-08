package mocks

import (
	"io"
	"net/http"
	"strings"
)

type MockClient struct {
	MockGet func(url string) (*http.Response, error)
}

func (c *MockClient) Get(url string) (*http.Response, error) {
	return c.MockGet(url)
}

func GetSuccess() *MockClient {
	return &MockClient{
		MockGet: func(url string) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`OK`)),
			}, nil
		},
	}
}
