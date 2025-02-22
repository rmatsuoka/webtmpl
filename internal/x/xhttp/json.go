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
	buf, err := json.Marshal(reqbody)
	if err != nil {
		return nil, fmt.Errorf("xhttp: new request with JSON: %w", err)
	}
	reader := bytes.NewReader(buf)
	req, err := NewRequest(ctx, method, url, reader, options...)
	if err != nil {
		return nil, fmt.Errorf("xhttp: new request with JSON: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// If c is nill, use http.DefaultClient.
func DoJSON(c Client, resbody any, req *http.Request) error {
	if c == nil {
		c = http.DefaultClient
	}
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

func PostJSON(ctx context.Context, c Client, resbody any, url string, reqbody any, options ...RequestOption) error {
	req, err := NewRequestJSON(ctx, http.MethodPost, url, reqbody, options...)
	if err != nil {
		return err
	}
	return DoJSON(c, resbody, req)
}

func GetJSON(ctx context.Context, c Client, resbody any, url string, options ...RequestOption) error {
	req, err := NewRequest(ctx, http.MethodGet, url, nil, options...)
	if err != nil {
		return err
	}
	return DoJSON(c, resbody, req)
}
