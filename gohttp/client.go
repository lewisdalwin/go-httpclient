// Filename: gohttp/client.go
package gohttp

import "net/http"

// create the interface
type HttpClient interface {
	// Get must be provided with the url of the resource, 
	// any headers that are to included in the request.
	// There is no body in a Get request
	Get(url string, headers http.Header) (*http.Response, error)
	// Post must be provided with the url of the enpoint to post to, 
	// any headers that are to be included in the request and any data that
	// should be used in the request
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	// Put must be provided with the url of the enpoint to update, 
	// any headers that are to be included in the request and any data that
	// should be used in the request
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	// Patch must be provided with the url of the resource to update, 
	// any headers that are to be included in the request and any data that
	// should be used in the request
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	// Delete must be provided with the url of the resource to delete, 
	// any headers that are to be included in the request.
	// There is no body in a Delete request
	Delete(url string, headers http.Header) (*http.Response, error)
}
// create a private type based on the interface
// needed to implement all the method signatures
type httpClient struct{}
// create a function to create a new instance 
// of the concrete type based on the interface
func New() HttpClient {
	client := &httpClient{}
	return client
}
// implement the methods on the type to satisfy the interface
func (c *httpClient) 	Get(url string, headers http.Header) (*http.Response, error){
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) 	Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) 	Put(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) 	Patch(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) 	Delete(url string, headers http.Header) (*http.Response, error){
	return c.do(http.MethodDelete, url, headers, nil)
}


