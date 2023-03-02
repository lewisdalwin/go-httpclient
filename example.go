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
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)
	return client
}

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
	// use the singleton client to make the request
	res, err := githubHttpClient.Get(url, headers)
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
