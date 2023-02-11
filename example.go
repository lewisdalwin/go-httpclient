// Filename: go-httpclient/example.go
package main

import (
	"fmt"
	"io"

	"github.com/lewisdalwin/go-httpclient/gohttp"
)

func main() {
	url := "https://api.github.com"
	client := gohttp.New()
	res, err := client.Get(url, nil)
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