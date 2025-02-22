package xhttp

import (
	"context"
	"io"
	"maps"
	"net/http"
)

type RequestOption func(*http.Request)

func WithHeader(h http.Header) RequestOption {
	return func(r *http.Request) {
		maps.Copy(r.Header, h)
	}
}

func NewRequest(ctx context.Context, method, url string, body io.Reader, options ...RequestOption) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		option(req)
	}
	return req, nil
}
