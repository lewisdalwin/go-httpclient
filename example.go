// Filename: go-httpclient/example.go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lewisdalwin/go-httpclient/gohttp"
)

func main() {
	url := "https://api.github.com"
	client := gohttp.New()
	// construct the headers
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	res, err := client.Get(url, headers)
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
