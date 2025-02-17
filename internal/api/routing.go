package api

import (
	"iter"
	"maps"
	"net/http"
)

func Handlers() iter.Seq2[string, http.Handler] {
	return maps.All(map[string]http.Handler{
		"GET  /api/count": http.HandlerFunc(count),
		"POST /api/count": http.HandlerFunc(countup),
	})
}
