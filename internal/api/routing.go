package api

import (
	"iter"
	"maps"
	"net/http"

	"github.com/rmatsuoka/webtmpl/internal/env"
	"github.com/rmatsuoka/webtmpl/internal/x/xhttp"
	"github.com/rmatsuoka/webtmpl/internal/x/xiter"
)

func Handlers() iter.Seq2[string, http.Handler] {
	handlers := maps.All(map[string]http.Handler{
		"GET  /api/count": http.HandlerFunc(count),
		"POST /api/count": http.HandlerFunc(countup),
	})

	return xiter.Map2(func(p string, h http.Handler) (string, http.Handler) {
		return p, xhttp.CSRF(env.APP_CSRF_ORIGINS, h)
	}, handlers)
}
