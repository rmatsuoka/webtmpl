package xhttp

import (
	"fmt"
	"net/http"
)

// NGStatusError is returned if a status code of Response is over 400.
type NGStatusError struct {
	// Response.Body is consumed and set Body.
	Response *http.Response
	Body     []byte
}

func (e *NGStatusError) Error() string {
	return fmt.Sprintf("xhttp: status code %d from %s %s: %.100s",
		e.Response.StatusCode, e.Response.Request.Method, e.Response.Request.URL.String(), e.Body)
}
