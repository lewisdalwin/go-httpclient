// Filename: go-httpclient/example.go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lewisdalwin/go-httpclient/gohttp"
)

// singeton (one instance)
var (
	httpClient = gohttp.New()
)

func main() {
	// reuse the client
	getData()
	getData()
	getData()
}

func getData() {
	url := "https://api.github.com"
	// construct the headers
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	// use the singleton client to make the request
	res, err := httpClient.Get(url, headers)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
