package http_client

import "net/http"

type HttpClientInterface interface {
	Get(url string) (*http.Response, error)
}
