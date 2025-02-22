package xhttp

import (
	"log/slog"
	"net/http"
	"net/url"
	"slices"
)

func CSRF(origins []string, h http.Handler) http.Handler {
	validRequest := func(req *http.Request) bool {
		origin := req.Header.Get(HeaderOrigin)
		return slices.Contains(origins, origin)
	}
	if len(origins) == 0 {
		validRequest = func(req *http.Request) bool {
			// Host Header does not include a `schema` field.
			origin, err := url.Parse(req.Header.Get(HeaderOrigin))
			if err != nil {
				return false
			}
			return origin.Host == req.Host
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet || req.Method == http.MethodHead {
			h.ServeHTTP(w, req)
			return
		}

		origin := req.Header.Get(HeaderOrigin)
		if origin == "" {
			http.Error(w, "No Origin header", http.StatusBadRequest)
			return
		}
		if !validRequest(req) {
			slog.ErrorContext(req.Context(), "detect CSRF attack",
				"origin", origin,
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
