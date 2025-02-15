package xhttp

import "net/http"

type Client struct {
	*http.Client
}

var DefaultClient = &Client{http.DefaultClient}
