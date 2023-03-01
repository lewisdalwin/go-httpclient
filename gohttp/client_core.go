// Filename: gohttp/client_core.go
package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	// create a new client
	client := http.Client{}
	// create a new request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}
	// TODO: set the headers
	for key, value := range headers {
		if len(value) > 0 {
			req.Header.Set(key, value[0])
		}
	}

	// issue the request
	return client.Do(req)

}
