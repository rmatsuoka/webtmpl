package xhttp

import (
	"maps"
	"net/http"
)

type RequestOption func(*http.Request)

func WithHeader(h http.Header) RequestOption {
	return func(r *http.Request) {
		maps.Copy(r.Header, h)
	}
}
