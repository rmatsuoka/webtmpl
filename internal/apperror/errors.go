package apperror

import (
	"errors"
	"net/http"
)

type appError struct {
	message    string
	statusCode int
}

func (e *appError) Error() string {
	return e.message
}

var errs = []*appError{}

func newError(message string, statusCode int) error {
	e := &appError{message, statusCode}
	errs = append(errs, e)
	return e
}

func lookup(err error) (*appError, bool) {
	for _, e := range errs {
		if errors.Is(err, e) {
			return e, true
		}
	}
	return nil, false
}

var (
	ErrNotFound = newError("not found", http.StatusNotFound)
)
