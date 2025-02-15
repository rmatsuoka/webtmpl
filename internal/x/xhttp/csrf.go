package xhttp

import (
	"net/http"
	"slices"
)

func CSRF(origins []string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodOptions {
			h.ServeHTTP(w, r)
			return
		}

		origin := r.Header.Get("Origin")
		if !slices.Contains(origins, origin) {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, r)
	})
}
