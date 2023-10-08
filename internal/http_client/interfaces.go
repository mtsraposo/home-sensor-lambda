package http_client

import "net/http"

type Client interface {
	Get(url string) (*http.Response, error)
}
