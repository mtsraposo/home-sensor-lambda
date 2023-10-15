package mocks

import (
	"io"
	"net/http"
	"strings"
)

type MockHttpClient struct {
	MockGet func(url string) (*http.Response, error)
}

func (c *MockHttpClient) Get(url string) (*http.Response, error) {
	return c.MockGet(url)
}

func HttpGetSuccess() *MockHttpClient {
	return &MockHttpClient{
		MockGet: func(url string) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`OK`)),
			}, nil
		},
	}
}
