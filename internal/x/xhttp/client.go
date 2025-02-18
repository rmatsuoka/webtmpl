package xhttp

import "net/http"

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

var _ Client = (*http.Client)(nil)
