package api

import (
	"fmt"
	"iter"
	"maps"
	"net/http"
)

func Handlers() iter.Seq2[string, http.Handler] {
	return maps.All(map[string]http.Handler{
		"GET /api/hello": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello")
		}),
	})
}
