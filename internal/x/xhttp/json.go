package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewRequestJSON(ctx context.Context, method, url string, reqbody any, options ...RequestOption) (*http.Request, error) {
	var reader io.Reader
	if reqbody != nil {
		buf, err := json.Marshal(reqbody)
		if err != nil {
			return nil, fmt.Errorf("xhttp: new request with JSON: %w", err)
		}
		reader = bytes.NewReader(buf)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return nil, fmt.Errorf("xhttp: new request with JSON: %w", err)
	}

	for _, o := range options {
		o(req)
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) DoJSON(resbody any, req *http.Request) error {
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		io.Copy(io.Discard, res.Body)
		res.Body.Close()
	}()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return &NGStatusError{
			Response: res,
			Body:     buf,
		}
	}

	if resbody == nil {
		return nil
	}

	err = json.Unmarshal(buf, resbody)
	if err != nil {
		return fmt.Errorf("xhttp: parse response body as JSON: %w", err)
	}
	return nil
}

func (c *Client) PostJSON(ctx context.Context, resbody any, url string, reqbody any, options ...RequestOption) error {
	req, err := NewRequestJSON(ctx, http.MethodPost, url, reqbody, options...)
	if err != nil {
		return err
	}
	return c.DoJSON(resbody, req)
}

func PostJSON(ctx context.Context, resbody any, url string, reqbody any, options ...RequestOption) error {
	return DefaultClient.PostJSON(ctx, resbody, url, reqbody, options...)
}

func (c *Client) GetJSON(ctx context.Context, resbody any, url string, options ...RequestOption) error {
	req, err := NewRequestJSON(ctx, http.MethodGet, url, nil, options...)
	if err != nil {
		return err
	}
	return c.DoJSON(resbody, req)
}

func GetJSON(ctx context.Context, resbody any, url string, options ...RequestOption) error {
	return DefaultClient.GetJSON(ctx, resbody, url, options...)
}
