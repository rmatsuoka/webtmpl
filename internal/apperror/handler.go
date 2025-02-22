package apperror

import (
	"net/http"

	"github.com/rmatsuoka/webtmpl/internal/x/xhttp"
)

func WriteJSON(w http.ResponseWriter, err error) {
	var (
		statusCode = http.StatusInternalServerError
		message    = http.StatusText(statusCode)
	)

	e, ok := lookup(err)
	if ok {
		statusCode = e.statusCode
		message = e.message
	}
	xhttp.WriteJSON(w, statusCode, xhttp.ErrorJSON{Message: message})
}
