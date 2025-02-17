package xhttp

import (
	"encoding/json"
	"io"
	"net/http"
)

type ErrorJSON struct {
	Message string `json:"message"`
}

func JSONHandler[T any](h func(http.ResponseWriter, *http.Request, T)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		rc := http.MaxBytesReader(w, req.Body, 1<<20)
		b, err := io.ReadAll(rc)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, ErrorJSON{
				Message: err.Error(),
			})
			return
		}
		var body T
		if err = json.Unmarshal(b, &body); err != nil {
			WriteJSON(w, http.StatusBadRequest, ErrorJSON{
				Message: err.Error(),
			})
			return
		}
		h(w, req, body)
	})
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	buf, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(statusCode)
	w.Write(buf)
}
