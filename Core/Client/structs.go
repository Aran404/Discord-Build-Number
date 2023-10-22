package client

import (
	http "github.com/bogdanfinn/fhttp"

	tls_client "github.com/bogdanfinn/tls-client"
)

type Options struct {
	Settings []tls_client.HttpClientOption
}

type ClientOptions interface {
	SetTimeout(int)
	SetNewCookieJar()
	SetNotFollowRedirects()
	SetProxy(string)
	NewClient() (*tls_client.HttpClient, error)
}

type Client struct {
	tls_client.HttpClient
}

type RequestResponse struct {
	Error                error
	StatusCode           int
	Ok                   bool
	StatusCodeDefinition string
	Body                 []byte
	Json                 map[string]interface{}
	Request              *http.Response
}