// Filename: go-httpclient/example.go
package go_httpclient

import "github.com/lewisdalwin/go-httpclient/gohttp"

func exampleUsage() {
	client := gohttp.New()
	client.Get()

}