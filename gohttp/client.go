// Filename: gohttp/client.go
package gohttp

// create the interface
type HttpClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
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
func (c *httpClient) Get() {}
func (c *httpClient) Post() {}
func (c *httpClient) Put() {}
func (c *httpClient) Patch() {}
func (c *httpClient) Delete() {}


