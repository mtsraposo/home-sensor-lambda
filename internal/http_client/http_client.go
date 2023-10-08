package http_client

import (
	"errors"
	"io"
	"log"
	"net/http"
)

type Client interface {
	Get(url string) (*http.Response, error)
}

type ClientImplementation struct{}

func (c *ClientImplementation) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func Get(url string, client Client) (string, error) {
	response, err := client.Get(url)
	responseError := handleResponseError(response, err, url)
	if responseError != nil {
		return "", responseError
	}
	statusCodeError := handleStatusCode(response, url)
	if statusCodeError != nil {
		return "", statusCodeError
	}

	body, err := readBody(response)
	if err != nil {
		return "", err
	}

	return body, nil
}

func handleResponseError(response *http.Response, err error, url string) error {
	if err != nil {
		log.Printf("Failed to make a request to %s: %v", url, err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print("Response body closed.")
		}
	}(response.Body)
	return nil
}

func handleStatusCode(response *http.Response, url string) error {
	if response.StatusCode != http.StatusOK {
		statusCodeError := errors.New("received non-OK HTTP status: " + response.Status)
		log.Printf("Request to %s failed: %v", url, statusCodeError)
		return statusCodeError
	}
	return nil
}

func readBody(response *http.Response) (string, error) {
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return "", err
	}
	return string(bodyBytes), nil
}
