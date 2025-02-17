package xhttp

import (
	"log/slog"
	"net/http"
	"slices"
)

func CSRF(origins []string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet || req.Method == http.MethodOptions {
			h.ServeHTTP(w, req)
			return
		}

		origin := req.Header.Get(HeaderOrigin)
		if origin == "" {
			http.Error(w, "No Origin header", http.StatusBadRequest)
			return
		}
		if !slices.Contains(origins, origin) {
			slog.ErrorContext(req.Context(), "CSRF attach",
				"method", req.Method,
				"path", req.URL.RequestURI(),
				"remoteAddr", req.RemoteAddr,
			)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, req)
	})
}
